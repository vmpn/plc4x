/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
// package org.apache.plc4x.rust.modbus.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use heapless::Vec;
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;


#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUReadFileRecordResponseItemOptions {
}

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUReadFileRecordResponseItem {
    pub referenceType: u8,
    pub data: Vec<u8, 1024>
}

impl ModbusPDUReadFileRecordResponseItem {
    pub fn dataLength(&self) -> u8 {
        (self.data.len() + 1) as u8
    }
}

impl Message for ModbusPDUReadFileRecordResponseItem {
    type M = ModbusPDUReadFileRecordResponseItem;
    type P = ModbusPDUReadFileRecordResponseItemOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u8(self.dataLength())?;
        writer.write_u8(self.referenceType)?;
        writer.write_bytes(self.data.as_slice())?;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let dataLength = reader.read_u8()?;
        let referenceType = reader.read_u8()?;
        let data: Vec<_, 1024>  = Vec::new();
        let data_read = 0 as usize;
        // while data_read < DefaultBinaryTerm{a=DefaultVariableLiteral{name='dataLength', typeReference='null', args=null, index=null, child=null}, b=DefaultNumericLiteral{number=1}, operation='-'} {
            // do something
        // }
        Ok(Self::M {
            referenceType,
            data
        })
    }
}

