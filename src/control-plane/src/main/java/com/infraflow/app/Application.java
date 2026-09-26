package com.infraflow.app;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * Main entry point for the Spring Boot application.
 *
 * The @SpringBootApplication annotation enables:
 * - Auto-configuration: Spring Boot automatically configures beans based on
 *   the dependencies present in the classpath
 * - Component scanning: Spring scans this package and sub-packages for
 *   components, services, repositories, and controllers
 * - Configuration: This class can define @Bean methods to create custom beans
 */

@SpringBootApplication
public class Application {

    /**
     * The main method that bootstraps the Spring application.
     *
     * SpringApplication.run() does the following:
     * 1. Creates an ApplicationContext
     * 2. Registers a shutdown hook
     * 3. Refreshes the context (loading beans)
     * 4. Starts the embedded web server (if spring-boot-starter-web is present)
     *
     * @param args Com mand line arguments passed to the application
     */

    public static void main(String[] args){
        SpringApplication.run(Application.class, args);
    }
}