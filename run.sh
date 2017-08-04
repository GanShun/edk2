#!/bin/bash

OPTS="-M q35 -m 2048 -smp 2"
OPTS="$OPTS -drive if=pflash,format=raw,file=Build/OvmfX64/DEBUG_GCC48/FV/OVMF.fd"
#OPTS="$OPTS -drive if=pflash,format=raw,file=Build/OvmfX64/DEBUG_GCC48/FV/OVMFNew.fd"
#OPTS="$OPTS -drive if=pflash,format=raw,file=~/bios/ocp/board2.dxectomy.replaceshellcode.rom"
#OPTS="$OPTS -drive if=pflash,format=raw,file=~/bios/ocp/smihell.rom"
OPTS="$OPTS -L /usr/local/google/home/ganshun/qemu/pc-bios"
#OPTS="$OPTS -device qxl-vga"
OPTS="$OPTS -display curses"
OPTS="$OPTS -monitor /dev/pts/27"
OPTS="$OPTS -serial /dev/tty"
OPTS="$OPTS -net none"
OPTS="$OPTS -global isa-debugcon.iobase=0x402 -debugcon file:fedora.ovmf.log"
OPTS="$OPTS -global PIIX4_PM.disable_s3=0"
qemu-system-x86_64 $OPTS
reset
