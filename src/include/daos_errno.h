/**
 * (C) Copyright 2016-2020 Intel Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
 * The Government's rights to use, modify, reproduce, release, perform, display,
 * or disclose this software are subject to the terms of the Apache License as
 * provided in Contract No. B609815.
 * Any reproduction of computer software, computer software documentation, or
 * portions thereof marked with this legend must also reproduce the markings.
 */
/**
 * DAOS Error numbers
 */

#ifndef __DAOS_ERRNO_H__
#define __DAOS_ERRNO_H__

#if defined(__cplusplus)
extern "C" {
#endif

/** @addtogroup GURT
 * @{
 */

/*
 * This preprocessor machinery defines the errno values but also
 * enables the internal definition of d_errstr.  A new macro should
 * be defined for each non-contiguous range
 */

#define D_FOREACH_GURT_ERR(ACTION)					\
	/** no permission */						\
	ACTION(DER_NO_PERM,		(DER_ERR_GURT_BASE + 1))	\
	/** invalid handle */						\
	ACTION(DER_NO_HDL,		(DER_ERR_GURT_BASE + 2))	\
	/** invalid parameters */					\
	ACTION(DER_INVAL,		(DER_ERR_GURT_BASE + 3))	\
	/** entity already exists */					\
	ACTION(DER_EXIST,		(DER_ERR_GURT_BASE + 4))	\
	/** nonexistent entity */					\
	ACTION(DER_NONEXIST,		(DER_ERR_GURT_BASE + 5))	\
	/** unreachable node */						\
	ACTION(DER_UNREACH,		(DER_ERR_GURT_BASE + 6))	\
	/** no space on storage target */				\
	ACTION(DER_NOSPACE,		(DER_ERR_GURT_BASE + 7))	\
	/** already did sth */						\
	ACTION(DER_ALREADY,		(DER_ERR_GURT_BASE + 8))	\
	/** NO memory */						\
	ACTION(DER_NOMEM,		(DER_ERR_GURT_BASE + 9))	\
	/** Function not implemented */					\
	ACTION(DER_NOSYS,		(DER_ERR_GURT_BASE + 10))	\
	/** timed out */						\
	ACTION(DER_TIMEDOUT,		(DER_ERR_GURT_BASE + 11))	\
	/** Busy */							\
	ACTION(DER_BUSY,		(DER_ERR_GURT_BASE + 12))	\
	/** Try again */						\
	ACTION(DER_AGAIN,		(DER_ERR_GURT_BASE + 13))	\
	/** incompatible protocol */					\
	ACTION(DER_PROTO,		(DER_ERR_GURT_BASE + 14))	\
	/** not initialized */						\
	ACTION(DER_UNINIT,		(DER_ERR_GURT_BASE + 15))	\
	/** buffer too short (larger buffer needed) */			\
	ACTION(DER_TRUNC,		(DER_ERR_GURT_BASE + 16))	\
	/** data too long for defined data type or buffer size */	\
	ACTION(DER_OVERFLOW,		(DER_ERR_GURT_BASE + 17))	\
	/** operation canceled */					\
	ACTION(DER_CANCELED,		(DER_ERR_GURT_BASE + 18))	\
	/** Out-Of-Group or member list */				\
	ACTION(DER_OOG,			(DER_ERR_GURT_BASE + 19))	\
	/** transport layer mercury error */				\
	ACTION(DER_HG,			(DER_ERR_GURT_BASE + 20))	\
	/** RPC or protocol version not registered */			\
	ACTION(DER_UNREG,		(DER_ERR_GURT_BASE + 21))	\
	/** failed to generate an address string */			\
	ACTION(DER_ADDRSTR_GEN,		(DER_ERR_GURT_BASE + 22))	\
	/** PMIx layer error */						\
	ACTION(DER_PMIX,		(DER_ERR_GURT_BASE + 23))	\
	/** IV callback - cannot handle locally */			\
	ACTION(DER_IVCB_FORWARD,	(DER_ERR_GURT_BASE + 24))	\
	/** miscellaneous error */					\
	ACTION(DER_MISC,		(DER_ERR_GURT_BASE + 25))	\
	/** Bad path name */						\
	ACTION(DER_BADPATH,		(DER_ERR_GURT_BASE + 26))	\
	/** Not a directory */						\
	ACTION(DER_NOTDIR,		(DER_ERR_GURT_BASE + 27))	\
	/** corpc failed */						\
	ACTION(DER_CORPC_INCOMPLETE,	(DER_ERR_GURT_BASE + 28))	\
	/** no rank is subscribed to RAS */				\
	ACTION(DER_NO_RAS_RANK,		(DER_ERR_GURT_BASE + 29))	\
	/** service group not attached */				\
	ACTION(DER_NOTATTACH,		(DER_ERR_GURT_BASE + 30))	\
	/** version mismatch */						\
	ACTION(DER_MISMATCH,		(DER_ERR_GURT_BASE + 31))	\
	/** rank has been evicted */					\
	ACTION(DER_EVICTED,		(DER_ERR_GURT_BASE + 32))	\
	/** user-provided RPC handler didn't send reply back */		\
	ACTION(DER_NOREPLY,		(DER_ERR_GURT_BASE + 33))	\
	/** denial-of-service */					\
	ACTION(DER_DOS,			(DER_ERR_GURT_BASE + 34))       \
	/** Incorrect target for the RPC  */				\
	ACTION(DER_BAD_TARGET,		(DER_ERR_GURT_BASE + 35))	\
	/** Group versioning mismatch */				\
	ACTION(DER_GRPVER,		(DER_ERR_GURT_BASE + 36))
	/** TODO: add more error numbers */

#define D_FOREACH_DAOS_ERR(ACTION)					\
	/** Generic I/O error */					\
	ACTION(DER_IO,			(DER_ERR_DAOS_BASE + 1))	\
	/** Memory free error */					\
	ACTION(DER_FREE_MEM,		(DER_ERR_DAOS_BASE + 2))	\
	/** Entry not found */						\
	ACTION(DER_ENOENT,		(DER_ERR_DAOS_BASE + 3))	\
	/** Unknown object type */					\
	ACTION(DER_NOTYPE,		(DER_ERR_DAOS_BASE + 4))	\
	/** Unknown object schema */					\
	ACTION(DER_NOSCHEMA,		(DER_ERR_DAOS_BASE + 5))	\
	/** Object is not local */					\
	ACTION(DER_NOLOCAL,		(DER_ERR_DAOS_BASE + 6))	\
	/** stale pool map version */					\
	ACTION(DER_STALE,		(DER_ERR_DAOS_BASE + 7))	\
	/** Not service leader */					\
	ACTION(DER_NOTLEADER,		(DER_ERR_DAOS_BASE + 8))	\
	/** * Target create error */					\
	ACTION(DER_TGT_CREATE,		(DER_ERR_DAOS_BASE + 9))	\
	/** Epoch is read-only */					\
	ACTION(DER_EP_RO,		(DER_ERR_DAOS_BASE + 10))	\
	/** Epoch is too old, all data have been recycled */		\
	ACTION(DER_EP_OLD,		(DER_ERR_DAOS_BASE + 11))	\
	/** Key is too large */						\
	ACTION(DER_KEY2BIG,		(DER_ERR_DAOS_BASE + 12))	\
	/** Record is too large */					\
	ACTION(DER_REC2BIG,		(DER_ERR_DAOS_BASE + 13))	\
	/** IO buffers can't match object extents */			\
	ACTION(DER_IO_INVAL,		(DER_ERR_DAOS_BASE + 14))	\
	/** Event queue is busy */					\
	ACTION(DER_EQ_BUSY,		(DER_ERR_DAOS_BASE + 15))	\
	/** Domain of cluster component can't match */			\
	ACTION(DER_DOMAIN,		(DER_ERR_DAOS_BASE + 16))	\
	/** Service should shut down */					\
	ACTION(DER_SHUTDOWN,		(DER_ERR_DAOS_BASE + 17))	\
	/** Operation now in progress */				\
	ACTION(DER_INPROGRESS,		(DER_ERR_DAOS_BASE + 18))	\
	/** Not applicable. */						\
	ACTION(DER_NOTAPPLICABLE,	(DER_ERR_DAOS_BASE + 19))	\
	/** Not a service replica */					\
	ACTION(DER_NOTREPLICA,		(DER_ERR_DAOS_BASE + 20))	\
	/** Checksum error */						\
	ACTION(DER_CSUM,		(DER_ERR_DAOS_BASE + 21))	\
	/** Unsupported durable format */				\
	ACTION(DER_DF_INVAL,		(DER_ERR_DAOS_BASE + 22))	\
	/** Incompatible durable format version */			\
	ACTION(DER_DF_INCOMPT,		(DER_ERR_DAOS_BASE + 23))	\
	/** Record size error */					\
	ACTION(DER_REC_SIZE,		(DER_ERR_DAOS_BASE + 24))	\
	/** Used to indicate a transaction should restart */		\
	ACTION(DER_TX_RESTART,		(DER_ERR_DAOS_BASE + 25))	\
	/** Data lost or not recoverable */				\
	ACTION(DER_DATA_LOSS,		(DER_ERR_DAOS_BASE + 26))	\
	/** Operation canceled (non-crt) */				\
	ACTION(DER_OP_CANCELED,		(DER_ERR_DAOS_BASE + 27))	\
	/** TX is not committed, not sure whether committable or not */	\
	ACTION(DER_TX_BUSY,		(DER_ERR_DAOS_BASE + 28))

/** Defines the gurt error codes */
#define D_FOREACH_ERR_RANGE(ACTION)	\
	ACTION(GURT,	1000)		\
	ACTION(DAOS,	2000)

#define D_DEFINE_ERRNO(name, value) name = value,

#define D_DEFINE_RANGE_ERRNO(name, base)			\
	enum {							\
		DER_ERR_##name##_BASE		=	(base),	\
		D_FOREACH_##name##_ERR(D_DEFINE_ERRNO)		\
		DER_ERR_##name##_LIMIT,				\
	};

#define D_DEFINE_ERRSTR(name, value) #name,

#define D_DEFINE_RANGE_ERRSTR(name)				\
	static const char * const g_##name##_error_strings[] = {\
		D_FOREACH_##name##_ERR(D_DEFINE_ERRSTR)		\
	};

D_FOREACH_ERR_RANGE(D_DEFINE_RANGE_ERRNO)

/** Macro to register a range defined using D_DEFINE_RANGE macros */
#define D_REGISTER_RANGE(name)				\
	d_errno_register_range(DER_ERR_##name##_BASE,	\
			       DER_ERR_##name##_LIMIT,	\
			       g_##name##_error_strings)

/** Macro to deregister a range defined using D_DEFINE_RANGE macros */
#define D_DEREGISTER_RANGE(name)			\
	d_errno_deregister_range(DER_ERR_##name##_BASE)

#define DER_SUCCESS	0
#define DER_UNKNOWN	(DER_ERR_GURT_BASE + 500000)

/** Return a string associated with a registered gurt errno
 *
 * \param	rc[in]	The error code
 *
 * \return	String value for error code or DER_UNKNOWN
 */
const char *d_errstr(int rc);

/** Register error codes with gurt.  Use D_REGISTER_RANGE.
 *
 * \param	start[in]	Start of error range.  Actual errors start at
 *				\p start + 1
 * \param	end[in]		End of range.  All error codes should be less
 *				than \p end
 * \param	error_strings[in]	Array of strings.   Must be one per
 *					code in the range
 *
 * \return	0 on success, otherwise error code
 */
int d_errno_register_range(int start, int end,
			   const char * const *error_strings);

/** De-register error codes with gurt.  Use D_DEREGISTER_RANGE.
 *
 * \param	start[in]	Start of error range
 */
void d_errno_deregister_range(int start);


/** @}
 */

#ifndef DF_RC
#define DF_RC "%s(%d)"
#define DP_RC(rc) d_errstr(rc), rc
#endif /* DF_RC */

#if defined(__cplusplus)
}
#endif

#endif /*  __DAOS_ERRNO_H__ */
