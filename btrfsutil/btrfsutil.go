// GPL-2

// WARNING: This file has automatically been generated on Sun, 01 Nov 2020 21:26:34 EST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package btrfsutil

/*
#cgo LDFLAGS: -lbtrfsutil
#include <btrfsutil.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// Strerror function as declared in include/btrfsutil.h:79
func Strerror(err BtrfsError) string {
	cerr, cerrAllocMap := (C.enum_btrfs_util_error)(err), cgoAllocsUnknown
	__ret := C.btrfs_util_strerror(cerr)
	runtime.KeepAlive(cerrAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// Sync function as declared in include/btrfsutil.h:87
func Sync(path string) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	__ret := C.btrfs_util_sync(cpath)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SyncFd function as declared in include/btrfsutil.h:92
func SyncFd(fd int32) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	__ret := C.btrfs_util_sync_fd(cfd)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// StartSync function as declared in include/btrfsutil.h:103
func StartSync(path string, transid []uint) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	ctransid, ctransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&transid)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_start_sync(cpath, ctransid)
	runtime.KeepAlive(ctransidAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// StartSyncFd function as declared in include/btrfsutil.h:109
func StartSyncFd(fd int32, transid []uint) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	ctransid, ctransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&transid)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_start_sync_fd(cfd, ctransid)
	runtime.KeepAlive(ctransidAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// WaitSync function as declared in include/btrfsutil.h:118
func WaitSync(path string, transid uint) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	ctransid, ctransidAllocMap := (C.uint64_t)(transid), cgoAllocsUnknown
	__ret := C.btrfs_util_wait_sync(cpath, ctransid)
	runtime.KeepAlive(ctransidAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// WaitSyncFd function as declared in include/btrfsutil.h:123
func WaitSyncFd(fd int32, transid uint) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	ctransid, ctransidAllocMap := (C.uint64_t)(transid), cgoAllocsUnknown
	__ret := C.btrfs_util_wait_sync_fd(cfd, ctransid)
	runtime.KeepAlive(ctransidAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// IsSubvolume function as declared in include/btrfsutil.h:134
func IsSubvolume(path string) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	__ret := C.btrfs_util_is_subvolume(cpath)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// IsSubvolumeFd function as declared in include/btrfsutil.h:139
func IsSubvolumeFd(fd int32) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	__ret := C.btrfs_util_is_subvolume_fd(cfd)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumeId function as declared in include/btrfsutil.h:148
func SubvolumeId(path string, idRet []uint) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cidRet, cidRetAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&idRet)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_subvolume_id(cpath, cidRet)
	runtime.KeepAlive(cidRetAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumeIdFd function as declared in include/btrfsutil.h:154
func SubvolumeIdFd(fd int32, idRet []uint) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cidRet, cidRetAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&idRet)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_subvolume_id_fd(cfd, cidRet)
	runtime.KeepAlive(cidRetAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumePath function as declared in include/btrfsutil.h:168
func SubvolumePath(path string, id uint, pathRet [][]byte) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cid, cidAllocMap := (C.uint64_t)(id), cgoAllocsUnknown
	cpathRet, cpathRetAllocMap := unpackArgSSByte(pathRet)
	__ret := C.btrfs_util_subvolume_path(cpath, cid, cpathRet)
	packSSByte(pathRet, cpathRet)
	runtime.KeepAlive(cpathRetAllocMap)
	runtime.KeepAlive(cidAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumePathFd function as declared in include/btrfsutil.h:174
func SubvolumePathFd(fd int32, id uint, pathRet [][]byte) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cid, cidAllocMap := (C.uint64_t)(id), cgoAllocsUnknown
	cpathRet, cpathRetAllocMap := unpackArgSSByte(pathRet)
	__ret := C.btrfs_util_subvolume_path_fd(cfd, cid, cpathRet)
	packSSByte(pathRet, cpathRet)
	runtime.KeepAlive(cpathRetAllocMap)
	runtime.KeepAlive(cidAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumeInfo function as declared in include/btrfsutil.h:278
func SubvolumeInfo(path string, id uint, subvol []BtrfsSubvolumeInfo) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cid, cidAllocMap := (C.uint64_t)(id), cgoAllocsUnknown
	csubvol, csubvolAllocMap := unpackArgSBtrfsSubvolumeInfo(subvol)
	__ret := C.btrfs_util_subvolume_info(cpath, cid, csubvol)
	packSBtrfsSubvolumeInfo(subvol, csubvol)
	runtime.KeepAlive(csubvolAllocMap)
	runtime.KeepAlive(cidAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumeInfoFd function as declared in include/btrfsutil.h:284
func SubvolumeInfoFd(fd int32, id uint, subvol []BtrfsSubvolumeInfo) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cid, cidAllocMap := (C.uint64_t)(id), cgoAllocsUnknown
	csubvol, csubvolAllocMap := unpackArgSBtrfsSubvolumeInfo(subvol)
	__ret := C.btrfs_util_subvolume_info_fd(cfd, cid, csubvol)
	packSBtrfsSubvolumeInfo(subvol, csubvol)
	runtime.KeepAlive(csubvolAllocMap)
	runtime.KeepAlive(cidAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// GetSubvolumeReadOnly function as declared in include/btrfsutil.h:294
func GetSubvolumeReadOnly(path string, ret []bool) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cret, cretAllocMap := (*C._Bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ret)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_get_subvolume_read_only(cpath, cret)
	runtime.KeepAlive(cretAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// GetSubvolumeReadOnlyFd function as declared in include/btrfsutil.h:301
func GetSubvolumeReadOnlyFd(fd int32, ret []bool) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cret, cretAllocMap := (*C._Bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ret)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_get_subvolume_read_only_fd(cfd, cret)
	runtime.KeepAlive(cretAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SetSubvolumeReadOnly function as declared in include/btrfsutil.h:312
func SetSubvolumeReadOnly(path string, readOnly bool) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	creadOnly, creadOnlyAllocMap := (C._Bool)(readOnly), cgoAllocsUnknown
	__ret := C.btrfs_util_set_subvolume_read_only(cpath, creadOnly)
	runtime.KeepAlive(creadOnlyAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SetSubvolumeReadOnlyFd function as declared in include/btrfsutil.h:319
func SetSubvolumeReadOnlyFd(fd int32, readOnly bool) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	creadOnly, creadOnlyAllocMap := (C._Bool)(readOnly), cgoAllocsUnknown
	__ret := C.btrfs_util_set_subvolume_read_only_fd(cfd, creadOnly)
	runtime.KeepAlive(creadOnlyAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// GetDefaultSubvolume function as declared in include/btrfsutil.h:332
func GetDefaultSubvolume(path string, idRet []uint) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cidRet, cidRetAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&idRet)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_get_default_subvolume(cpath, cidRet)
	runtime.KeepAlive(cidRetAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// GetDefaultSubvolumeFd function as declared in include/btrfsutil.h:339
func GetDefaultSubvolumeFd(fd int32, idRet []uint) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cidRet, cidRetAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&idRet)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_get_default_subvolume_fd(cfd, cidRet)
	runtime.KeepAlive(cidRetAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SetDefaultSubvolume function as declared in include/btrfsutil.h:354
func SetDefaultSubvolume(path string, id uint) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cid, cidAllocMap := (C.uint64_t)(id), cgoAllocsUnknown
	__ret := C.btrfs_util_set_default_subvolume(cpath, cid)
	runtime.KeepAlive(cidAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SetDefaultSubvolumeFd function as declared in include/btrfsutil.h:361
func SetDefaultSubvolumeFd(fd int32, id uint) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cid, cidAllocMap := (C.uint64_t)(id), cgoAllocsUnknown
	__ret := C.btrfs_util_set_default_subvolume_fd(cfd, cid)
	runtime.KeepAlive(cidAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSubvolume function as declared in include/btrfsutil.h:377
func CreateSubvolume(path string, flags int32, asyncTransid []uint, qgroupInherit []BtrfsQgroupInherit) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	casyncTransid, casyncTransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&asyncTransid)).Data)), cgoAllocsUnknown
	cqgroupInherit, cqgroupInheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&qgroupInherit)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_create_subvolume(cpath, cflags, casyncTransid, cqgroupInherit)
	runtime.KeepAlive(cqgroupInheritAllocMap)
	runtime.KeepAlive(casyncTransidAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSubvolumeFd function as declared in include/btrfsutil.h:393
func CreateSubvolumeFd(parentFd int32, name string, flags int32, asyncTransid []uint, qgroupInherit []BtrfsQgroupInherit) BtrfsError {
	cparentFd, cparentFdAllocMap := (C.int)(parentFd), cgoAllocsUnknown
	cname, cnameAllocMap := unpackPCharString(name)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	casyncTransid, casyncTransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&asyncTransid)).Data)), cgoAllocsUnknown
	cqgroupInherit, cqgroupInheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&qgroupInherit)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_create_subvolume_fd(cparentFd, cname, cflags, casyncTransid, cqgroupInherit)
	runtime.KeepAlive(cqgroupInheritAllocMap)
	runtime.KeepAlive(casyncTransidAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cnameAllocMap)
	runtime.KeepAlive(cparentFdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSnapshot function as declared in include/btrfsutil.h:428
func CreateSnapshot(source string, path string, flags int32, asyncTransid []uint, qgroupInherit []BtrfsQgroupInherit) BtrfsError {
	csource, csourceAllocMap := unpackPCharString(source)
	cpath, cpathAllocMap := unpackPCharString(path)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	casyncTransid, casyncTransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&asyncTransid)).Data)), cgoAllocsUnknown
	cqgroupInherit, cqgroupInheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&qgroupInherit)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_create_snapshot(csource, cpath, cflags, casyncTransid, cqgroupInherit)
	runtime.KeepAlive(cqgroupInheritAllocMap)
	runtime.KeepAlive(casyncTransidAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	runtime.KeepAlive(csourceAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSnapshotFd function as declared in include/btrfsutil.h:436
func CreateSnapshotFd(fd int32, path string, flags int32, asyncTransid []uint, qgroupInherit []BtrfsQgroupInherit) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cpath, cpathAllocMap := unpackPCharString(path)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	casyncTransid, casyncTransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&asyncTransid)).Data)), cgoAllocsUnknown
	cqgroupInherit, cqgroupInheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&qgroupInherit)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_create_snapshot_fd(cfd, cpath, cflags, casyncTransid, cqgroupInherit)
	runtime.KeepAlive(cqgroupInheritAllocMap)
	runtime.KeepAlive(casyncTransidAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSnapshotFd2 function as declared in include/btrfsutil.h:452
func CreateSnapshotFd2(fd int32, parentFd int32, name string, flags int32, asyncTransid []uint, qgroupInherit []BtrfsQgroupInherit) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cparentFd, cparentFdAllocMap := (C.int)(parentFd), cgoAllocsUnknown
	cname, cnameAllocMap := unpackPCharString(name)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	casyncTransid, casyncTransidAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&asyncTransid)).Data)), cgoAllocsUnknown
	cqgroupInherit, cqgroupInheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&qgroupInherit)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_create_snapshot_fd2(cfd, cparentFd, cname, cflags, casyncTransid, cqgroupInherit)
	runtime.KeepAlive(cqgroupInheritAllocMap)
	runtime.KeepAlive(casyncTransidAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cnameAllocMap)
	runtime.KeepAlive(cparentFdAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// DeleteSubvolume function as declared in include/btrfsutil.h:476
func DeleteSubvolume(path string, flags int32) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	__ret := C.btrfs_util_delete_subvolume(cpath, cflags)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// DeleteSubvolumeFd function as declared in include/btrfsutil.h:487
func DeleteSubvolumeFd(parentFd int32, name string, flags int32) BtrfsError {
	cparentFd, cparentFdAllocMap := (C.int)(parentFd), cgoAllocsUnknown
	cname, cnameAllocMap := unpackPCharString(name)
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	__ret := C.btrfs_util_delete_subvolume_fd(cparentFd, cname, cflags)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(cnameAllocMap)
	runtime.KeepAlive(cparentFdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSubvolumeIterator function as declared in include/btrfsutil.h:524
func CreateSubvolumeIterator(path string, top uint, flags int32, ret [][]BtrfsSubvolumeIterator) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	ctop, ctopAllocMap := (C.uint64_t)(top), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	cret, cretAllocMap := unpackArgSSBtrfsSubvolumeIterator(ret)
	__ret := C.btrfs_util_create_subvolume_iterator(cpath, ctop, cflags, cret)
	packSSBtrfsSubvolumeIterator(ret, cret)
	runtime.KeepAlive(cretAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(ctopAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateSubvolumeIteratorFd function as declared in include/btrfsutil.h:533
func CreateSubvolumeIteratorFd(fd int32, top uint, flags int32, ret [][]BtrfsSubvolumeIterator) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	ctop, ctopAllocMap := (C.uint64_t)(top), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	cret, cretAllocMap := unpackArgSSBtrfsSubvolumeIterator(ret)
	__ret := C.btrfs_util_create_subvolume_iterator_fd(cfd, ctop, cflags, cret)
	packSSBtrfsSubvolumeIterator(ret, cret)
	runtime.KeepAlive(cretAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(ctopAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// DestroySubvolumeIterator function as declared in include/btrfsutil.h:543
func DestroySubvolumeIterator(iter []BtrfsSubvolumeIterator) {
	citer, citerAllocMap := (*C.struct_btrfs_util_subvolume_iterator)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&iter)).Data)), cgoAllocsUnknown
	C.btrfs_util_destroy_subvolume_iterator(citer)
	runtime.KeepAlive(citerAllocMap)
}

// SubvolumeIteratorFd function as declared in include/btrfsutil.h:556
func SubvolumeIteratorFd(iter []BtrfsSubvolumeIterator) int32 {
	citer, citerAllocMap := (*C.struct_btrfs_util_subvolume_iterator)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&iter)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_subvolume_iterator_fd(citer)
	runtime.KeepAlive(citerAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SubvolumeIteratorNext function as declared in include/btrfsutil.h:573
func SubvolumeIteratorNext(iter []BtrfsSubvolumeIterator, pathRet [][]byte, idRet []uint) BtrfsError {
	citer, citerAllocMap := (*C.struct_btrfs_util_subvolume_iterator)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&iter)).Data)), cgoAllocsUnknown
	cpathRet, cpathRetAllocMap := unpackArgSSByte(pathRet)
	cidRet, cidRetAllocMap := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&idRet)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_subvolume_iterator_next(citer, cpathRet, cidRet)
	runtime.KeepAlive(cidRetAllocMap)
	packSSByte(pathRet, cpathRet)
	runtime.KeepAlive(cpathRetAllocMap)
	runtime.KeepAlive(citerAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// SubvolumeIteratorNextInfo function as declared in include/btrfsutil.h:592
func SubvolumeIteratorNextInfo(iter []BtrfsSubvolumeIterator, pathRet [][]byte, subvol []BtrfsSubvolumeInfo) BtrfsError {
	citer, citerAllocMap := (*C.struct_btrfs_util_subvolume_iterator)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&iter)).Data)), cgoAllocsUnknown
	cpathRet, cpathRetAllocMap := unpackArgSSByte(pathRet)
	csubvol, csubvolAllocMap := unpackArgSBtrfsSubvolumeInfo(subvol)
	__ret := C.btrfs_util_subvolume_iterator_next_info(citer, cpathRet, csubvol)
	packSBtrfsSubvolumeInfo(subvol, csubvol)
	runtime.KeepAlive(csubvolAllocMap)
	packSSByte(pathRet, cpathRet)
	runtime.KeepAlive(cpathRetAllocMap)
	runtime.KeepAlive(citerAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// DeletedSubvolumes function as declared in include/btrfsutil.h:607
func DeletedSubvolumes(path string, ids [][]uint, n []uint) BtrfsError {
	cpath, cpathAllocMap := unpackPCharString(path)
	cids, cidsAllocMap := unpackArgSSUUint(ids)
	cn, cnAllocMap := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&n)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_deleted_subvolumes(cpath, cids, cn)
	runtime.KeepAlive(cnAllocMap)
	packSSUUint(ids, cids)
	runtime.KeepAlive(cidsAllocMap)
	runtime.KeepAlive(cpathAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// DeletedSubvolumesFd function as declared in include/btrfsutil.h:614
func DeletedSubvolumesFd(fd int32, ids [][]uint, n []uint) BtrfsError {
	cfd, cfdAllocMap := (C.int)(fd), cgoAllocsUnknown
	cids, cidsAllocMap := unpackArgSSUUint(ids)
	cn, cnAllocMap := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&n)).Data)), cgoAllocsUnknown
	__ret := C.btrfs_util_deleted_subvolumes_fd(cfd, cids, cn)
	runtime.KeepAlive(cnAllocMap)
	packSSUUint(ids, cids)
	runtime.KeepAlive(cidsAllocMap)
	runtime.KeepAlive(cfdAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// CreateQgroupInherit function as declared in include/btrfsutil.h:628
func CreateQgroupInherit(flags int32, ret [][]BtrfsQgroupInherit) BtrfsError {
	cflags, cflagsAllocMap := (C.int)(flags), cgoAllocsUnknown
	cret, cretAllocMap := unpackArgSSBtrfsQgroupInherit(ret)
	__ret := C.btrfs_util_create_qgroup_inherit(cflags, cret)
	packSSBtrfsQgroupInherit(ret, cret)
	runtime.KeepAlive(cretAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// DestroyQgroupInherit function as declared in include/btrfsutil.h:636
func DestroyQgroupInherit(inherit []BtrfsQgroupInherit) {
	cinherit, cinheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inherit)).Data)), cgoAllocsUnknown
	C.btrfs_util_destroy_qgroup_inherit(cinherit)
	runtime.KeepAlive(cinheritAllocMap)
}

// QgroupInheritAddGroup function as declared in include/btrfsutil.h:646
func QgroupInheritAddGroup(inherit [][]BtrfsQgroupInherit, qgroupid uint) BtrfsError {
	cinherit, cinheritAllocMap := unpackArgSSBtrfsQgroupInherit(inherit)
	cqgroupid, cqgroupidAllocMap := (C.uint64_t)(qgroupid), cgoAllocsUnknown
	__ret := C.btrfs_util_qgroup_inherit_add_group(cinherit, cqgroupid)
	runtime.KeepAlive(cqgroupidAllocMap)
	packSSBtrfsQgroupInherit(inherit, cinherit)
	runtime.KeepAlive(cinheritAllocMap)
	__v := (BtrfsError)(__ret)
	return __v
}

// QgroupInheritGetGroups function as declared in include/btrfsutil.h:656
func QgroupInheritGetGroups(inherit []BtrfsQgroupInherit, groups [][]uint, n []uint) {
	cinherit, cinheritAllocMap := (*C.struct_btrfs_util_qgroup_inherit)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inherit)).Data)), cgoAllocsUnknown
	cgroups, cgroupsAllocMap := unpackArgSSUUint(groups)
	cn, cnAllocMap := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&n)).Data)), cgoAllocsUnknown
	C.btrfs_util_qgroup_inherit_get_groups(cinherit, cgroups, cn)
	runtime.KeepAlive(cnAllocMap)
	packSSUUint(groups, cgroups)
	runtime.KeepAlive(cgroupsAllocMap)
	runtime.KeepAlive(cinheritAllocMap)
}