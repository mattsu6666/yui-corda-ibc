package jp.datachain.corda.ibc.clients.lcp

import com.google.protobuf.ByteString
import ibc.core.client.v1.Client
import ibc.core.commitment.v1.Commitment

data class StateCommitment(
    val prefix: Commitment.MerklePrefix,
    val path: String,
    val value: ByteArray,
    val height: Client.Height,
    val stateId: ByteArray
) {

    companion object {
        fun rlpDecode(rlp: ByteString): StateCommitment? {
            TODO()
        }
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as StateCommitment

        if (prefix != other.prefix) return false
        if (path != other.path) return false
        if (!value.contentEquals(other.value)) return false
        if (height != other.height) return false
        if (!stateId.contentEquals(other.stateId)) return false

        return true
    }

    override fun hashCode(): Int {
        var result = prefix.hashCode()
        result = 31 * result + path.hashCode()
        result = 31 * result + value.contentHashCode()
        result = 31 * result + height.hashCode()
        result = 31 * result + stateId.contentHashCode()
        return result
    }
}
