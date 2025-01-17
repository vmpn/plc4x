/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.codegen.ast;

public class BinaryExpression extends Expression {

    private final Node left;
    private final Node right;
    private final Operation op;

    BinaryExpression(TypeDefinition type, Node left, Node right, Operation op) {
        super(type);
        this.left = left;
        this.right = right;
        this.op = op;
    }

    public Node getLeft() {
        return left;
    }

    public Node getRight() {
        return right;
    }

    public Operation getOp() {
        return op;
    }

    @Override
    public <T> T accept(NodeVisitor<T> visitor) {
        return null;
    }

    @Override
    public void write(Generator writer) {
        writer.generate(this);
    }

    public enum Operation {
        EQ,
        NEQ,
        GT,
        LT,
        PLUS;
    }
}
