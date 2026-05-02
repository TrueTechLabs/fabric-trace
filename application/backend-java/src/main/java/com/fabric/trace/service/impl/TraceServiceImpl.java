package com.fabric.trace.service.impl;

import cn.hutool.core.io.FileUtil;
import cn.hutool.core.util.StrUtil;
import com.fabric.trace.common.Result;
import com.fabric.trace.service.FabricService;
import com.fabric.trace.service.TraceService;
import com.fabric.trace.util.SnowflakeUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.util.ArrayList;
import java.util.List;

@Service
public class TraceServiceImpl implements TraceService {

    @Autowired
    private FabricService fabricService;

    @Value("${file.upload-path}")
    private String uploadPath;

    @Value("${file.image-path}")
    private String imagePath;

    @Override
    public Result uplink(String userId, String traceabilityCode,
                         String arg1, String arg2, String arg3, String arg4, String arg5,
                         MultipartFile file) throws Exception {
        String userType = fabricService.chaincodeQuery("GetUserType", userId);

        List<String> args = new ArrayList<>();
        args.add(userId);

        if ("种植户".equals(userType)) {
            String farmerTraceCode = SnowflakeUtil.generateId().substring(1);
            args.add(farmerTraceCode);
            traceabilityCode = farmerTraceCode;
        } else {
            if (StrUtil.isBlank(traceabilityCode) || traceabilityCode.length() != 18) {
                return Result.error("请检查溯源码是否正确!!");
            }
            String res = fabricService.chaincodeQuery("GetFruitInfo", traceabilityCode);
            if (StrUtil.isBlank(res)) {
                return Result.error("请检查溯源码是否正确!!");
            }
            args.add(traceabilityCode);
        }

        args.add(arg1);
        args.add(arg2);
        args.add(arg3);
        args.add(arg4);
        args.add(arg5);

        if (file != null && !file.isEmpty()) {
            String originalFilename = file.getOriginalFilename();
            String ext = SnowflakeUtil.getFileExt(originalFilename);

            File tempDir = new File(uploadPath);
            if (!tempDir.exists()) {
                tempDir.mkdirs();
            }
            File tempFile = new File(uploadPath + originalFilename);
            file.transferTo(tempFile);

            String sha256 = SnowflakeUtil.calculateFileSHA256(tempFile.getAbsolutePath());
            String imageFilename = sha256 + "." + ext;
            File imageDir = new File(imagePath);
            if (!imageDir.exists()) {
                imageDir.mkdirs();
            }
            FileUtil.move(tempFile, new File(imagePath + imageFilename), true);

            args.add(imageFilename);
        } else {
            args.add("");
        }

        args.add("");

        String txid = fabricService.chaincodeInvoke("Uplink", args);

        return Result.success("uplink success")
                .txid(txid)
                .traceabilityCode(traceabilityCode);
    }

    @Override
    public Result getFruitInfo(String traceabilityCode) throws Exception {
        String res = fabricService.chaincodeQuery("GetFruitInfo", traceabilityCode);
        return Result.success("query success").data(res);
    }

    @Override
    public Result getFruitList(String userId) throws Exception {
        String res = fabricService.chaincodeQuery("GetFruitList", userId);
        return Result.success("query success").data(res);
    }

    @Override
    public Result getAllFruitInfo() throws Exception {
        String res = fabricService.chaincodeQuery("GetAllFruitInfo", "");
        return Result.success("query success").data(res);
    }

    @Override
    public Result getFruitHistory(String traceabilityCode) throws Exception {
        String res = fabricService.chaincodeQuery("GetFruitHistory", traceabilityCode);
        return Result.success("query success").data(res);
    }
}
