package jp.datachain.corda.cross.contracts

import cross.core.initiator.Msgs
import cross.core.initiator.Msgs.MsgInitiateTx
import jp.datachain.corda.cross.state.InitiateTxState
import jp.datachain.corda.cross.state.TxAuthState
import net.corda.core.contracts.StateRef
import net.corda.core.contracts.UniqueIdentifier
import net.corda.core.crypto.SecureHash
import net.corda.testing.core.TestIdentity
import net.corda.testing.node.MockServices
import net.corda.testing.node.internal.setDriverSerialization
import net.corda.testing.node.ledger
import org.junit.Test

class CrossContractTests {
    private val alice = TestIdentity.fresh("Alice")
    private val validator1 = TestIdentity.fresh("Validator1")
    private val validator2 = TestIdentity.fresh("Validator2")

    private val ledgerServices = MockServices(listOf("jp.datachain.corda.cross"))

    @Test
    fun initiateTxMustHasTwoOutputs() {
        val baseId = StateRef(SecureHash.zeroHash, 0)
        val txId = UniqueIdentifier()
        val participants = listOf(alice, validator1, validator2)
        val msg = MsgInitiateTx.newBuilder()
                .setChainId(baseId.toString())
                .build()

        val txAuthState = TxAuthState.of(txId, participants.map{it.party}.toSet(), emptySet())
        val initiateTxState = InitiateTxState.of(txId, participants.map{it.party}.toSet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg)

        setDriverSerialization(ClassLoader.getSystemClassLoader()).use {
            ledgerServices.ledger {
                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    command(participants.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    fails()
                }

                transaction {
                    output(CrossContract::class.qualifiedName!!, initiateTxState)
                    command(participants.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    fails()
                }

                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    output(CrossContract::class.qualifiedName!!, initiateTxState)
                    command(participants.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    verifies()
                }
            }
        }
    }

    @Test
    fun initiateTxMustHasSameValueAsBaseIdAndChainId() {
        val baseId = StateRef(SecureHash.zeroHash, 0)
        val txId = UniqueIdentifier()
        val participants = listOf(alice, validator1, validator2)
        val msg1 = MsgInitiateTx.newBuilder()
                .setChainId("another_chain_id")
                .build()
        val msg2 = MsgInitiateTx.newBuilder()
                .setChainId(baseId.toString())
                .build()

        val txAuthState = TxAuthState.of(txId, participants.map{it.party}.toSet(), emptySet())
        val initiateTxState1 = InitiateTxState.of(txId, participants.map{it.party}.toSet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg1)
        val initiateTxState2 = InitiateTxState.of(txId, participants.map{it.party}.toSet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg2)

        setDriverSerialization(ClassLoader.getSystemClassLoader()).use {
            ledgerServices.ledger {
                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    output(CrossContract::class.qualifiedName!!, initiateTxState1)
                    command(participants.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    fails()
                }

                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    output(CrossContract::class.qualifiedName!!, initiateTxState2)
                    command(participants.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    verifies()
                }
            }
        }
    }

    @Test
    fun initiateTxMustHasSameTxIdAndParticipants() {
        val baseId = StateRef(SecureHash.zeroHash, 0)
        val txId1 = UniqueIdentifier()
        val txId2 = UniqueIdentifier()
        val participants1 = listOf(alice, validator1, validator2)
        val participants2 = listOf(alice, validator1)
        val msg = MsgInitiateTx.newBuilder()
                .setChainId(baseId.toString())
                .build()

        val txAuthState = TxAuthState.of(txId1, participants1.map{it.party}.toSet(), emptySet())
        val initiateTxState1 = InitiateTxState.of(txId1, participants1.map{it.party}.toSet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg)
        val initiateTxState2 = InitiateTxState.of(txId1, participants2.map{it.party}.toSet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg)
        val initiateTxState3 = InitiateTxState.of(txId2, participants1.map{it.party}.toSet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg)

        setDriverSerialization(ClassLoader.getSystemClassLoader()).use {
            ledgerServices.ledger {
                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    output(CrossContract::class.qualifiedName!!, initiateTxState2)
                    command(participants1.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    fails()
                }

                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    output(CrossContract::class.qualifiedName!!, initiateTxState3)
                    command(participants1.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    fails()
                }

                transaction {
                    output(CrossContract::class.qualifiedName!!, txAuthState)
                    output(CrossContract::class.qualifiedName!!, initiateTxState1)
                    command(participants1.map { it.party.owningKey }, CrossContract.Commands.InitiateTx())
                    verifies()
                }
            }
        }
    }
}
