package com.fabric.trace.service;

import java.util.List;

public interface FabricService {

    String chaincodeQuery(String fcn, String arg) throws Exception;

    String chaincodeInvoke(String fcn, List<String> args) throws Exception;
}
