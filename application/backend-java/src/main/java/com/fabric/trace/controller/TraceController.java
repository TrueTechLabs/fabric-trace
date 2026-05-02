package com.fabric.trace.controller;

import com.fabric.trace.common.Result;
import com.fabric.trace.service.TraceService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.FileSystemResource;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import java.io.File;
import java.net.URLEncoder;

@RestController
public class TraceController {

    @Autowired
    private TraceService traceService;

    @Value("${file.image-path}")
    private String imagePath;

    @PostMapping("/uplink")
    public Result uplink(@RequestParam(required = false) String traceability_code,
                         @RequestParam String arg1,
                         @RequestParam String arg2,
                         @RequestParam String arg3,
                         @RequestParam String arg4,
                         @RequestParam String arg5,
                         @RequestPart(required = false) MultipartFile file,
                         HttpServletRequest request) throws Exception {
        String userId = (String) request.getAttribute("userID");
        return traceService.uplink(userId, traceability_code, arg1, arg2, arg3, arg4, arg5, file);
    }

    @PostMapping("/getFruitInfo")
    public Result getFruitInfo(@RequestParam String traceability_code) throws Exception {
        return traceService.getFruitInfo(traceability_code);
    }

    @PostMapping("/getFruitList")
    public Result getFruitList(HttpServletRequest request) throws Exception {
        String userId = (String) request.getAttribute("userID");
        return traceService.getFruitList(userId);
    }

    @PostMapping("/getAllFruitInfo")
    public Result getAllFruitInfo() throws Exception {
        return traceService.getAllFruitInfo();
    }

    @PostMapping("/getFruitHistory")
    public Result getFruitHistory(@RequestParam String traceability_code) throws Exception {
        return traceService.getFruitHistory(traceability_code);
    }

    @GetMapping("/getImg/{filename}")
    public ResponseEntity<Resource> getImg(@PathVariable String filename) throws Exception {
        File file = new File(imagePath + filename);
        if (!file.exists()) {
            return ResponseEntity.notFound().build();
        }
        FileSystemResource resource = new FileSystemResource(file);
        return ResponseEntity.ok()
                .header(HttpHeaders.CONTENT_DISPOSITION,
                        "inline; filename=\"" + URLEncoder.encode(filename, "UTF-8") + "\"")
                .contentType(MediaType.IMAGE_JPEG)
                .body(resource);
    }

    @GetMapping("/ping")
    public Result ping() {
        return Result.success("pong");
    }
}
