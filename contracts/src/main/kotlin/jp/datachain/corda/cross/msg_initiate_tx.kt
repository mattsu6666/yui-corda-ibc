package jp.datachain.corda.cross

import cross.core.initiator.Msgs

fun Msgs.MsgInitiateTx.requiredSigners(): Set<cross.core.auth.Types.Account> {
    return contractTransactionsList.flatMap {
        it.signersList
    }.toSet()
}
