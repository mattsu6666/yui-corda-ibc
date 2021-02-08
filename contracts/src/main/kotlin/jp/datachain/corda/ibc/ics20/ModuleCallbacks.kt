package jp.datachain.corda.ibc.ics20

import com.google.protobuf.ByteString
import ibc.applications.transfer.v1.Transfer
import ibc.core.channel.v1.ChannelOuterClass
import jp.datachain.corda.ibc.ics24.Identifier
import jp.datachain.corda.ibc.ics26.Context
import jp.datachain.corda.ibc.ics26.ModuleCallbacks

class ModuleCallbacks: ModuleCallbacks {
    override fun onRecvPacket(ctx: Context, packet: ChannelOuterClass.Packet): ChannelOuterClass.Acknowledgement {
        val data = Transfer.FungibleTokenPacketData.parseFrom(packet.data)
        val denom = Denom(data.denom)
        val amount = Amount(data.amount)
        val receiver = Address(data.receiver)

        val ackBuilder = ChannelOuterClass.Acknowledgement.newBuilder()
        val source = denom.hasPrefix(Identifier(packet.sourcePort), Identifier(packet.sourceChannel))
        val bank = ctx.getInput<Bank>()
        if (source) {
            try {
                val unprefixedDenom = denom.removePrefix()
                ctx.addOutput(bank.unlock(receiver, unprefixedDenom, amount))
                ackBuilder.result = ByteString.copyFrom(ByteArray(1){1})
            } catch (e: IllegalArgumentException) {
                ctx.addOutput(bank.copy())
                ackBuilder.error = e.message!!
            }
        } else {
            try {
                val prefixedDenom = denom.addPrefix(Identifier(packet.destinationPort), Identifier(packet.destinationChannel))
                ctx.addOutput(bank
                        .recordDenom(prefixedDenom)
                        .mint(receiver, prefixedDenom.ibcDenom, amount))
                ackBuilder.result = ByteString.copyFrom(ByteArray(1){1})
            } catch (e: IllegalArgumentException) {
                ctx.addOutput(bank.copy())
                ackBuilder.error = e.message!!
            }
        }
        return ackBuilder.build()
    }

    override fun onAcknowledgePacket(ctx: Context, packet: ChannelOuterClass.Packet, acknowledgement: ChannelOuterClass.Acknowledgement) {
        when (acknowledgement.responseCase) {
            ChannelOuterClass.Acknowledgement.ResponseCase.RESULT ->
                ctx.addOutput(ctx.getInput<Bank>().copy())
            ChannelOuterClass.Acknowledgement.ResponseCase.ERROR ->
                refundTokens(ctx, packet)
            else -> throw java.lang.IllegalArgumentException()
        }
    }

    private fun refundTokens(ctx: Context, packet: ChannelOuterClass.Packet) {
        val data = Transfer.FungibleTokenPacketData.parseFrom(packet.data)
        val denom = Denom(data.denom)
        val amount = Amount(data.amount)
        val sender = Address(data.sender)

        val source = !denom.hasPrefix(Identifier(packet.sourcePort), Identifier(packet.sourceChannel))
        val bank = ctx.getInput<Bank>()
        if (source) {
            ctx.addOutput(bank.unlock(sender, denom, amount))
        } else {
            ctx.addOutput(bank.mint(sender, denom.ibcDenom, amount))
        }
    }
}