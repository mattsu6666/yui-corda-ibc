package jp.datachain.corda.cross.serialization

import cross.core.initiator.State
import net.corda.core.serialization.SerializationCustomSerializer
import net.corda.core.utilities.OpaqueBytes

class InitiateTxStateSerializer: SerializationCustomSerializer<State.InitiateTxState, InitiateTxStateSerializer.Proxy> {
    data class Proxy(val serialized: OpaqueBytes)

    override fun toProxy(obj: State.InitiateTxState) = Proxy(OpaqueBytes(obj.toByteArray() + 0))
    override fun fromProxy(proxy: Proxy) = State.InitiateTxState.parseFrom(proxy.serialized.bytes.let { it.copyOf(it.size - 1) })!!
}
