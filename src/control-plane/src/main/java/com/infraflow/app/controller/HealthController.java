package com.infraflow.app.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * A simple REST controller demonstrating basic Spring MVC concepts.
 *
 * @RestController combines @Controller and @ResponseBody, meaning:
 * - This class handles HTTP requests
 * - Return values are automatically serialized to JSON
 */
@RestController
public class HealthController {

    @GetMapping("/health")
    public String health() {
        return "Springboot up and running!";
    }

}