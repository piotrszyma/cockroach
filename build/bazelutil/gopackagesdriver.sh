#!/bin/bash
exec bazelisk run -- @io_bazel_rules_go//go/tools/gopackagesdriver "${@}"
