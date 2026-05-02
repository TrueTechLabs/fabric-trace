package com.fabric.trace.config;

import io.grpc.ManagedChannel;
import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;
import io.grpc.netty.shaded.io.netty.handler.ssl.util.InsecureTrustManagerFactory;
import org.hyperledger.fabric.client.Contract;
import org.hyperledger.fabric.client.Gateway;
import org.hyperledger.fabric.client.Network;
import org.hyperledger.fabric.client.identity.Identities;
import org.hyperledger.fabric.client.identity.Identity;
import org.hyperledger.fabric.client.identity.Signer;
import org.hyperledger.fabric.client.identity.Signers;
import org.hyperledger.fabric.client.identity.X509Identity;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.PrivateKey;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.concurrent.TimeUnit;

@Configuration
public class FabricConfig {

    @Value("${fabric.msp-id}")
    private String mspId;

    @Value("${fabric.peer-endpoint}")
    private String peerEndpoint;

    @Value("${fabric.tls-cert-path}")
    private String tlsCertPath;

    @Value("${fabric.cert-path}")
    private String certPath;

    @Value("${fabric.key-path}")
    private String keyPath;

    @Value("${fabric.channel-name}")
    private String channelName;

    @Value("${fabric.chaincode-name}")
    private String chaincodeName;

    @Bean
    public Contract fabricContract() throws Exception {
        // TLS证书
        X509Certificate tlsCert = readCertificate(tlsCertPath);
        SslContext sslContext = GrpcSslContexts.forClient()
                .trustManager(tlsCert)
                .build();

        ManagedChannel grpcChannel = io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder
                .forTarget(peerEndpoint)
                .sslContext(sslContext)
                .overrideAuthority("peer0.org1.example.com")
                .build();

        // 用户身份
        X509Certificate certificate = readCertificate(certPath);
        Identity identity = new X509Identity(mspId, certificate);

        // 私钥签名
        PrivateKey privateKey = readPrivateKey(keyPath);
        Signer signer = Signers.newPrivateKeySigner(privateKey);

        // 构建Gateway
        Gateway gateway = Gateway.newInstance()
                .identity(identity)
                .signer(signer)
                .connection(grpcChannel)
                .evaluateOptions(opts -> opts.withDeadlineAfter(5, TimeUnit.SECONDS))
                .endorseOptions(opts -> opts.withDeadlineAfter(15, TimeUnit.SECONDS))
                .submitOptions(opts -> opts.withDeadlineAfter(5, TimeUnit.SECONDS))
                .commitStatusOptions(opts -> opts.withDeadlineAfter(60, TimeUnit.SECONDS))
                .connect();

        Network network = gateway.getNetwork(channelName);
        return network.getContract(chaincodeName);
    }

    private X509Certificate readCertificate(String path) throws IOException, CertificateException {
        String pem = Files.readString(Paths.get(path));
        return Identities.readX509Certificate(pem);
    }

    private PrivateKey readPrivateKey(String keyDir) throws IOException, InvalidKeyException {
        Path dir = Paths.get(keyDir);
        try (var stream = Files.list(dir)) {
            Path firstFile = stream.findFirst()
                    .orElseThrow(() -> new IOException("私钥目录为空: " + keyDir));
            String pem = Files.readString(firstFile);
            return Identities.readPrivateKey(pem);
        }
    }
}
