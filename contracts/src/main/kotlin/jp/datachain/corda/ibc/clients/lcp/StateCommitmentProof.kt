package jp.datachain.corda.ibc.clients.lcp

import com.google.protobuf.ByteString

data class StateCommitmentProof(
    val commitmentBytes: ByteArray,
    val signer: String,
    val signature: ByteArray
) {

    companion object {
        fun decodeRlp(rlp: ByteString): StateCommitmentProof? {
            TODO()
        }
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as StateCommitmentProof

        if (!commitmentBytes.contentEquals(other.commitmentBytes)) return false
        if (signer != other.signer) return false
        if (!signature.contentEquals(other.signature)) return false

        return true
    }

    override fun hashCode(): Int {
        var result = commitmentBytes.contentHashCode()
        result = 31 * result + signer.hashCode()
        result = 31 * result + signature.contentHashCode()
        return result
    }
}
