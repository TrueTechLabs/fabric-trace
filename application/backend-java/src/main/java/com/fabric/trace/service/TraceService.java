package com.fabric.trace.service;

import com.fabric.trace.common.Result;
import org.springframework.web.multipart.MultipartFile;

public interface TraceService {

    Result uplink(String userId, String traceabilityCode,
                  String arg1, String arg2, String arg3, String arg4, String arg5,
                  MultipartFile file) throws Exception;

    Result getFruitInfo(String traceabilityCode) throws Exception;

    Result getFruitList(String userId) throws Exception;

    Result getAllFruitInfo() throws Exception;

    Result getFruitHistory(String traceabilityCode) throws Exception;
}
