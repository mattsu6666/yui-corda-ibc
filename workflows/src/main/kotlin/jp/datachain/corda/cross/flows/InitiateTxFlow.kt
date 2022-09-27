package jp.datachain.corda.cross.flows

import co.paralleluniverse.fibers.Suspendable
import cross.core.initiator.Msgs
import jp.datachain.corda.cross.contracts.CrossContract
import jp.datachain.corda.cross.requiredSigners
import jp.datachain.corda.cross.state.InitiateTxState
import jp.datachain.corda.cross.state.TxAuthState
import net.corda.core.contracts.StateRef
import net.corda.core.contracts.UniqueIdentifier
import net.corda.core.flows.*
import net.corda.core.identity.Party
import net.corda.core.transactions.SignedTransaction
import net.corda.core.transactions.TransactionBuilder

@InitiatingFlow
@StartableByRPC
class InitiateTxFlow(private val baseId: StateRef, private val msg: Msgs.MsgInitiateTx) : FlowLogic<SignedTransaction>() {
    @Suspendable
    override fun call(): SignedTransaction {
        val notary = serviceHub.networkMapCache.notaryIdentities.single()
        val validators: Set<Party> = emptySet() // TODO: implement
        val txId = UniqueIdentifier()

        val participants = validators + ourIdentity
        val txAuthState = TxAuthState
                .of(txId, participants, msg.requiredSigners())
                .sign(msg.signersList)
        val its = InitiateTxState.of(txId, participants, baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg)
        val initiateTxState = if (txAuthState.completed()) {
            its.toVerified()
        } else {
            its
        }

        val txBuilder = TransactionBuilder(notary)
                .addOutputState(txAuthState)
                .addOutputState(initiateTxState)
                .addCommand(
                        CrossContract.Commands.InitiateTx(),
                        participants.map { it.owningKey }
                )

        txBuilder.verify(serviceHub)
        val signedTx = serviceHub.signInitialTransaction(txBuilder)

        // Sharing states with validators
        val otherSessions = validators.map { initiateFlow(it) }
        return subFlow(FinalityFlow(signedTx, otherSessions))
    }
}

@InitiatedBy(InitiateTxFlow::class)
class InitiateTxResponderFlow(private val counterPartySession: FlowSession) : FlowLogic<Unit>() {
    @Suspendable
    override fun call() {
        val stx = subFlow(ReceiveFinalityFlow(counterPartySession))
        println(stx)
    }
}
