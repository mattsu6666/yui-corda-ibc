package jp.datachain.corda.cross.contracts

import jp.datachain.corda.cross.state.InitiateTxState
import jp.datachain.corda.cross.state.TxAuthState
import net.corda.core.contracts.CommandData
import net.corda.core.contracts.Contract
import net.corda.core.contracts.requireSingleCommand
import net.corda.core.contracts.requireThat
import net.corda.core.transactions.LedgerTransaction

class CrossContract : Contract {

    @Throws(IllegalArgumentException::class)
    override fun verify(tx: LedgerTransaction) {
        val command = tx.commands.requireSingleCommand<Commands>()
        when (command.value) {
            is Commands.InitiateTx -> requireThat {
                "No state should be consumed" using (tx.inputs.isEmpty())
                "Exactly two states should be created" using (tx.outputs.size == 2)

                // TODO: Check if skipping timeout height and timestamp validation should be skipped
                val newInitiateTxState = tx.outputsOfType<InitiateTxState>().single()
                "BaseId should be equal to the ChainId of MsgInitiateTx" using
                        (newInitiateTxState.baseId.toString() == newInitiateTxState.inner.msg.chainId)

                val newTxAuthState = tx.outputsOfType<TxAuthState>().single()
                "InitiateTxState and TxAuthState have the same txId" using
                        (newInitiateTxState.linearId == newTxAuthState.linearId)
                "InitiateTxState and TxAuthState have the same participants" using
                        (newInitiateTxState.participants == newTxAuthState.participants)

                "All participants should sign the transaction" using
                        (command.signers.toSet() == newInitiateTxState.participants.map { it.owningKey }.toSet())
            }
        }
    }

    interface Commands : CommandData {
        class InitiateTx : Commands
    }
}
