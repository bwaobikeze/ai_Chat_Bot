package com.ppuu444.h2memdb.model;

import org.springframework.data.annotation.Id;

public record Executions(@Id Integer id, Integer hash, String result) {
}
