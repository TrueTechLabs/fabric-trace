package com.fabric.trace.interceptor;

import cn.hutool.core.util.StrUtil;
import com.fabric.trace.util.JwtUtil;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.jsonwebtoken.Claims;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@Component
public class JwtAuthInterceptor implements HandlerInterceptor {

    @Autowired
    private JwtUtil jwtUtil;

    @Autowired
    private ObjectMapper objectMapper;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if ("OPTIONS".equalsIgnoreCase(request.getMethod())) {
            return true;
        }

        String authHeader = request.getHeader("Authorization");
        if (StrUtil.isBlank(authHeader)) {
            writeError(response, 401, "请求未携带token，无权限访问");
            return false;
        }

        // 兼容 "Bearer xxx" 和直接发 token 两种格式
        String token = authHeader;
        if (authHeader.startsWith("Bearer ")) {
            token = authHeader.substring(7);
        }

        try {
            Claims claims = jwtUtil.parseToken(token);
            request.setAttribute("userID", claims.get("userID", String.class));
        } catch (Exception e) {
            writeError(response, 401, "请求未携带token，无权限访问");
            return false;
        }

        return true;
    }

    private void writeError(HttpServletResponse response, int code, String msg) throws Exception {
        response.setStatus(200);
        response.setContentType("application/json;charset=UTF-8");
        Map<String, Object> result = new HashMap<>();
        result.put("code", code);
        result.put("msg", msg);
        response.getWriter().write(objectMapper.writeValueAsString(result));
    }
}
