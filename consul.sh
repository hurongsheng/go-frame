#!/bin/bash

consul agent -dev -enable-script-checks -config-dir=./consul.d
