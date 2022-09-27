package jp.datachain.corda.cross.state

import cross.core.auth.Types
import jp.datachain.corda.cross.contracts.CrossContract
import net.corda.core.contracts.BelongsToContract
import net.corda.core.contracts.LinearState
import net.corda.core.contracts.UniqueIdentifier
import net.corda.core.identity.AbstractParty
import net.corda.core.identity.Party
import net.corda.core.serialization.ConstructorForDeserialization

@BelongsToContract(CrossContract::class)
data class TxAuthState @ConstructorForDeserialization constructor(
        override val linearId: UniqueIdentifier,
        override val participants: List<AbstractParty>,
        val inner: Types.TxAuthState
) : LinearState {

    fun sign(signers: List<Types.Account>): TxAuthState {
        val newSigners = inner.remainingSignersList.toSet() - signers.toSet()
        val newState = Types.TxAuthState
                .newBuilder(inner)
                .clearRemainingSigners()
                .addAllRemainingSigners(newSigners.toList())
                .build()
        return TxAuthState(linearId, participants, newState)
    }

    fun completed(): Boolean = inner.remainingSignersList.isEmpty()

    companion object {
        fun of(txId: UniqueIdentifier, participants: Set<Party>, required: Set<Types.Account>): TxAuthState {
            val state = Types.TxAuthState.newBuilder().addAllRemainingSigners(required.toList()).build()
            return TxAuthState(txId, participants.toList(), state)
        }
    }
}
