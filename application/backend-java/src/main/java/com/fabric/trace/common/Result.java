package com.fabric.trace.common;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Data;

@Data
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Result {

    private int code;
    private String message;
    private Object data;
    private String jwt;
    private String txid;
    private String traceability_code;
    private String userType;
    private String username;

    public static Result success(String message) {
        Result r = new Result();
        r.code = 200;
        r.message = message;
        return r;
    }

    public static Result error(String message) {
        Result r = new Result();
        r.code = 500;
        r.message = message;
        return r;
    }

    public Result data(Object data) {
        this.data = data;
        return this;
    }

    public Result jwt(String jwt) {
        this.jwt = jwt;
        return this;
    }

    public Result txid(String txid) {
        this.txid = txid;
        return this;
    }

    public Result traceabilityCode(String code) {
        this.traceability_code = code;
        return this;
    }

    public Result userType(String userType) {
        this.userType = userType;
        return this;
    }

    public Result username(String username) {
        this.username = username;
        return this;
    }
}
