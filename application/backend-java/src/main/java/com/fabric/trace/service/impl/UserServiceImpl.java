package com.fabric.trace.service.impl;

import cn.hutool.core.util.StrUtil;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.fabric.trace.common.Result;
import com.fabric.trace.entity.User;
import com.fabric.trace.mapper.UserMapper;
import com.fabric.trace.service.FabricService;
import com.fabric.trace.service.UserService;
import com.fabric.trace.util.SnowflakeUtil;
import com.fabric.trace.util.JwtUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Arrays;

@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private UserMapper userMapper;

    @Autowired
    private FabricService fabricService;

    @Autowired
    private JwtUtil jwtUtil;

    @Override
    public Result register(String username, String password, String userType) throws Exception {
        // 检查用户名是否已存在
        Long count = userMapper.selectCount(
                new LambdaQueryWrapper<User>().eq(User::getUsername, username));
        if (count > 0) {
            return Result.error("用户名已存在");
        }

        // 生成用户ID并存入MySQL
        User user = new User();
        user.setUserId(SnowflakeUtil.generateId());
        user.setUsername(username);
        user.setPassword(SnowflakeUtil.encryptByMD5(password));
        user.setRealInfo(SnowflakeUtil.encryptByMD5(username));
        userMapper.insert(user);

        // 注册到区块链
        String txid = fabricService.chaincodeInvoke("RegisterUser",
                Arrays.asList(user.getUserId(), userType, user.getRealInfo()));

        return Result.success("register success").txid(txid);
    }

    @Override
    public Result login(String username, String password) throws Exception {
        // 查询用户
        User user = userMapper.selectOne(
                new LambdaQueryWrapper<User>().eq(User::getUsername, username));
        if (user == null) {
            return Result.error("没有找到该用户");
        }

        // 校验密码
        if (!SnowflakeUtil.encryptByMD5(password).equals(user.getPassword())) {
            return Result.error("密码错误");
        }

        // 从区块链获取用户类型
        String userType = fabricService.chaincodeQuery("GetUserType", user.getUserId());
        if (StrUtil.isBlank(userType)) {
            return Result.error("login failed: 获取用户类型失败");
        }

        // 生成JWT
        String jwt = jwtUtil.generateToken(user.getUserId(), userType);

        return Result.success("login success").jwt(jwt);
    }

    @Override
    public Result logout() {
        return Result.success("logout success");
    }

    @Override
    public Result getInfo(String userId) throws Exception {
        // 从区块链获取用户类型
        String userType = fabricService.chaincodeQuery("GetUserType", userId);

        // 从MySQL获取用户名
        User user = userMapper.selectById(userId);
        if (user == null) {
            return Result.error("get user type failed");
        }

        return Result.success("get user type success")
                .userType(userType)
                .username(user.getUsername());
    }
}
