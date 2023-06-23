// Package directoryReader provides functions for reading the contents of a drive or directory
package directoryReader


import (
    "log"
    "syscall"
)


// GetDrives returns a slive of runes representing the names of the drives present on the system
// This function uses syscall to call the GetLogicalDrives function from the kernel32 library
func GetDrives() (drives []rune) {
    // Loads the kernel32 library for the GetLogicalDrives function
    kernel32, err :=  syscall.LoadLibrary("kernel32.dll")
    if err != nil {
        log.Fatal(err)
    }
    defer syscall.FreeLibrary(kernel32)

    // Gets the address of the GetLogicalDrives function
    getLogicalDrives, err := syscall.GetProcAddress(kernel32, "GetLogicalDrives")
    if err != nil {
        log.Fatal(err)
    }

    // Calls the GetLogicalDrives function
    // This returns a bit map of the drives where the lowest bit is drive A
    r, _, err := syscall.SyscallN(uintptr(getLogicalDrives), 0, 0, 0, 0)
    if err != nil {
        log.Fatal(err)
    }
    bitMap := uint32(r)

    // Converts the bit map to a slice of runes representing the names of the drives
    for i := 0; i < 26; i++ {
        // If the lowest bit is 1, then the drive exists
        // Same as bitMap % 00..001 == 00..001
        if bitMap & 1 == 1 {
            drives = append(drives, rune('A'+i))
        }
        bitMap >>= 1
    }

    return
}