#!/usr/bin/env python3

import os

archs = {
    "linux": ["amd64", "arm64"],
    "windows": ["amd64"],
    "darwin": ["amd64", "arm64"],
}

for os_name, arch_list in archs.items():
    for arch_name in arch_list:
        output_name = f"dirsplitter_{os_name}_{arch_name}"

        if os_name == "windows":
            output_name += ".exe"

        print(f"Building {output_name}...")

        os.system(
            f"GOOS={os_name} GOARCH={arch_name} go build -o ./build/{output_name} dirsplitter.go"
        )
