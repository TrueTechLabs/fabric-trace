package com.fabric.trace.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.fabric.trace.entity.User;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface UserMapper extends BaseMapper<User> {
}
