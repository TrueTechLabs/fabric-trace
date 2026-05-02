package com.fabric.trace.service;

import com.fabric.trace.common.Result;

public interface UserService {

    Result register(String username, String password, String userType) throws Exception;

    Result login(String username, String password) throws Exception;

    Result logout();

    Result getInfo(String userId) throws Exception;
}
