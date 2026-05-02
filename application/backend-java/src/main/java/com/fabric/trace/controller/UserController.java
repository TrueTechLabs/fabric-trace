package com.fabric.trace.controller;

import com.fabric.trace.common.Result;
import com.fabric.trace.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;

@RestController
public class UserController {

    @Autowired
    private UserService userService;

    @PostMapping("/register")
    public Result register(@RequestParam String username,
                           @RequestParam String password,
                           @RequestParam String userType) throws Exception {
        return userService.register(username, password, userType);
    }

    @PostMapping("/login")
    public Result login(@RequestParam String username,
                        @RequestParam String password) throws Exception {
        return userService.login(username, password);
    }

    @PostMapping("/logout")
    public Result logout() {
        return userService.logout();
    }

    @PostMapping("/getInfo")
    public Result getInfo(HttpServletRequest request) throws Exception {
        String userId = (String) request.getAttribute("userID");
        return userService.getInfo(userId);
    }
}
