package jp.datachain.corda.cross.serialization

import cross.core.auth.Types
import net.corda.core.serialization.SerializationCustomSerializer
import net.corda.core.utilities.OpaqueBytes

class TxAuthStateSerializer: SerializationCustomSerializer<Types.TxAuthState, TxAuthStateSerializer.Proxy> {
    data class Proxy(val serialized: OpaqueBytes)

    override fun toProxy(obj: Types.TxAuthState) = Proxy(OpaqueBytes(obj.toByteArray() + 0))
    override fun fromProxy(proxy: Proxy) = Types.TxAuthState.parseFrom(proxy.serialized.bytes.let { it.copyOf(it.size - 1) })!!
}
