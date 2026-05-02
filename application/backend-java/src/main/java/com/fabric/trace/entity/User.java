package com.fabric.trace.entity;

import com.baomidou.mybatisplus.annotation.TableField;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;

@Data
@TableName("users")
public class User {

    @TableId("user_id")
    private String userId;

    @TableField("username")
    private String username;

    @TableField("password")
    private String password;

    @TableField("realInfo")
    private String realInfo;
}
