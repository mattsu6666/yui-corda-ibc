package jp.datachain.corda.ibc.ics23

import net.corda.core.serialization.CordaSerializable
import net.corda.core.utilities.OpaqueBytes

@CordaSerializable
class CommitmentPrefix(bytes: ByteArray): OpaqueBytes(bytes)