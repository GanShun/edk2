#!/bin/bash
OPTS="-M q35 -m 2048 -smp 2"
OPTS="$OPTS -drive if=pflash,format=raw,file="
f=`echo Build/OvmfX64/*/FV/OVMF.fd`
OPTS="$OPTS$f"
OPTS="$OPTS -display none"
OPTS="$OPTS -monitor $MONITOR"
OPTS="$OPTS -serial /dev/tty"
OPTS="$OPTS -global isa-debugcon.iobase=0x402 -debugcon file:fedora.ovmf.log"
OPTS="$OPTS -net none"
echo $OPTS
qemu-system-x86_64 $OPTS $*
reset
