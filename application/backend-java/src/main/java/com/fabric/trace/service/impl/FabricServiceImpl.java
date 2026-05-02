package com.fabric.trace.service.impl;

import com.fabric.trace.service.FabricService;
import org.hyperledger.fabric.client.Contract;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.nio.charset.StandardCharsets;
import java.util.List;

@Service
public class FabricServiceImpl implements FabricService {

    @Autowired
    private Contract contract;

    @Override
    public String chaincodeQuery(String fcn, String arg) throws Exception {
        byte[] result;
        if (arg == null || arg.isEmpty()) {
            result = contract.evaluateTransaction(fcn);
        } else {
            result = contract.evaluateTransaction(fcn, arg);
        }
        return new String(result, StandardCharsets.UTF_8);
    }

    @Override
    public String chaincodeInvoke(String fcn, List<String> args) throws Exception {
        var proposal = contract.newProposal(fcn)
                .addArguments(args.toArray(new String[0]))
                .build();
        var transaction = proposal.endorse().submitAsync();

        var status = transaction.getStatus();
        if (!status.isSuccessful()) {
            throw new RuntimeException("交易提交失败: " + status.getCode());
        }

        return transaction.getTransactionId();
    }
}
