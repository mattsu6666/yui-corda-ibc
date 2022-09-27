package jp.datachain.corda.cross.state

import cross.core.initiator.Msgs
import net.corda.core.contracts.StateRef
import net.corda.core.contracts.UniqueIdentifier
import net.corda.core.crypto.SecureHash
import org.junit.Test

class InitiateTxStateTest {

    @Test
    fun toVerified() {
        val baseId = StateRef(SecureHash.zeroHash, 0)
        val txId = UniqueIdentifier()
        val msg = Msgs.MsgInitiateTx.newBuilder()
                .setChainId(baseId.toString())
                .build()
        val initiateTxState = InitiateTxState.of(txId, emptySet(), baseId, Msgs.InitiateTxStatus.INITIATE_TX_STATUS_PENDING, msg)

        val verified = initiateTxState.toVerified()
        assert(verified.inner.status == Msgs.InitiateTxStatus.INITIATE_TX_STATUS_VERIFIED)
    }
}
