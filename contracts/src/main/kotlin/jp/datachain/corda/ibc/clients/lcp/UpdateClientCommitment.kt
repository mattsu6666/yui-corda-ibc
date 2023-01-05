package jp.datachain.corda.ibc.clients.lcp

import com.google.protobuf.ByteString
import ibc.core.client.v1.Client

data class UpdateClientCommitment(
    val prevStateId: ByteArray,
    val newStateId: ByteArray,
    val newState: ByteArray,
    val prevHeight: Client.Height,
    val newHeight: Client.Height,
    val timestamp: Long,
    val validationParams: ByteArray
) {

    companion object {
        fun decodeRlp(rlp: ByteString): UpdateClientCommitment? {
            TODO()
        }
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as UpdateClientCommitment

        if (!prevStateId.contentEquals(other.prevStateId)) return false
        if (!newStateId.contentEquals(other.newStateId)) return false
        if (!newState.contentEquals(other.newState)) return false
        if (prevHeight != other.prevHeight) return false
        if (newHeight != other.newHeight) return false
        if (timestamp != other.timestamp) return false
        if (!validationParams.contentEquals(other.validationParams)) return false

        return true
    }

    override fun hashCode(): Int {
        var result = prevStateId.contentHashCode()
        result = 31 * result + newStateId.contentHashCode()
        result = 31 * result + newState.contentHashCode()
        result = 31 * result + prevHeight.hashCode()
        result = 31 * result + newHeight.hashCode()
        result = 31 * result + timestamp.hashCode()
        result = 31 * result + validationParams.contentHashCode()
        return result
    }
}
