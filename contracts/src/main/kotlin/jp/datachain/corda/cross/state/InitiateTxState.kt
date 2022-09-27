package jp.datachain.corda.cross.state

import cross.core.initiator.Msgs
import jp.datachain.corda.cross.contracts.CrossContract
import net.corda.core.contracts.BelongsToContract
import net.corda.core.contracts.LinearState
import net.corda.core.contracts.StateRef
import net.corda.core.contracts.UniqueIdentifier
import net.corda.core.identity.AbstractParty
import net.corda.core.identity.Party
import net.corda.core.serialization.ConstructorForDeserialization

@BelongsToContract(CrossContract::class)
class InitiateTxState @ConstructorForDeserialization constructor(
        override val linearId: UniqueIdentifier,
        override val participants: List<AbstractParty>,
        val baseId: StateRef, // TODO: Needed?
        val inner: cross.core.initiator.State.InitiateTxState
) : LinearState {

        fun toVerified(): InitiateTxState {
                return InitiateTxState(
                        this.linearId,
                        this.participants,
                        this.baseId,
                        cross.core.initiator.State.InitiateTxState.newBuilder(inner)
                                .setStatus(Msgs.InitiateTxStatus.INITIATE_TX_STATUS_VERIFIED)
                                .build()
                )
        }

        companion object {
                fun of(
                        txId: UniqueIdentifier,
                        participants: Set<Party>,
                        baseId: StateRef,
                        status: Msgs.InitiateTxStatus,
                        msg: Msgs.MsgInitiateTx): InitiateTxState {
                        val state = cross.core.initiator.State.InitiateTxState.newBuilder()
                                .setStatus(status)
                                .setMsg(msg)
                                .build()
                        return InitiateTxState(txId, participants.toList(), baseId, state)
                }
        }
}
