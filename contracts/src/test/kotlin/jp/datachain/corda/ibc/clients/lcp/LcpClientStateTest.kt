package jp.datachain.corda.ibc.clients.lcp

import com.r3.conclave.common.OpaqueBytes
import com.r3.conclave.common.internal.attestation.EpidAttestation
import org.junit.Test
import java.security.cert.CertificateFactory

class LcpClientStateTest {

    @Test
    fun test() {
        val factory = CertificateFactory.getInstance("X.509")
        EpidAttestation(OpaqueBytes.parse(""), OpaqueBytes.parse(""), factory.generateCertPath(listOf()))
    }
}
