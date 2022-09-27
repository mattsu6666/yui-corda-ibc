package jp.datachain.corda.cross.state

import com.google.protobuf.ByteString
import cross.core.auth.Types
import net.corda.core.contracts.UniqueIdentifier
import org.junit.Test
import kotlin.test.assertTrue

class TxAuthStateTest {

    @Test
    fun sign() {
        val txId = UniqueIdentifier()
        val account1 = Types.Account.newBuilder().setId(ByteString.copyFromUtf8("account1")).build()
        val account2 = Types.Account.newBuilder().setId(ByteString.copyFromUtf8("account2")).build()
        val accounts = setOf(account1, account2)
        val txAuthState = TxAuthState.of(txId, emptySet(), accounts)

        val signed = txAuthState.sign(listOf(account1))
        assert(signed.inner.remainingSignersList == listOf(account2))
    }

    @Test
    fun completed() {
        val txId = UniqueIdentifier()
        val account1 = Types.Account.newBuilder().setId(ByteString.copyFromUtf8("account1")).build()
        val account2 = Types.Account.newBuilder().setId(ByteString.copyFromUtf8("account2")).build()
        val accounts = setOf(account1, account2)
        val txAuthState = TxAuthState.of(txId, emptySet(), accounts)

        val signed = txAuthState.sign(listOf(account1, account2))
        assertTrue(signed.completed())
    }
}
