// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package disk

// ReReadPartitionTable forces the kernel to re-read the partition table
// on the disk.
//
// It is done via an ioctl call with request as BLKRRPART.
func (d *Disk) ReReadPartitionTable() error {
	return nil
}
