package com.fabric.trace.util;

import cn.hutool.core.io.FileUtil;
import cn.hutool.core.lang.Snowflake;
import cn.hutool.crypto.digest.DigestUtil;

import java.io.File;

public class SnowflakeUtil {

    private static final Snowflake SNOWFLAKE = cn.hutool.core.util.IdUtil.getSnowflake(1, 1);

    public static String generateId() {
        return SNOWFLAKE.nextIdStr();
    }

    public static String encryptByMD5(String str) {
        return DigestUtil.md5Hex(str);
    }

    public static String calculateFileSHA256(String filePath) {
        return DigestUtil.sha256Hex(new File(filePath));
    }

    public static String getFileExt(String filename) {
        return FileUtil.extName(filename);
    }
}
