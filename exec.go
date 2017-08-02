package libgosu

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Exec attempts to apply the userspec to the current thread before calling
// syscall.Exec which itself invokes the execve(2) system call. Param `userspec`
// should be in the form: "<user or uid>" or "<user or uid>:<group or guid>"
//
// NOTE: On a successful call to syscall.Exec this function will never return.
// See http://man7.org/linux/man-pages/man2/execve.2.html:
//
// "execve() does not return on success, and the text, initialized data,
//  uninitialized data (bss), and stack of the calling process are
//  overwritten according to the contents of the newly loaded program."
func Exec(userspec string, cmdwargs []string) error {
	// clear HOME so that SetupUser will set it
	os.Unsetenv("HOME")
	if err := SetupUser(userspec); err != nil {
		return errors.New(fmt.Sprintf("failed switching to %q: %v", userspec, err))
	}
	name, err := exec.LookPath(cmdwargs[0])
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}
	if err = syscall.Exec(name, cmdwargs, os.Environ()); err != nil {
		return errors.New(fmt.Sprintf("exec failed: %v", err))
	}
	return nil
}
