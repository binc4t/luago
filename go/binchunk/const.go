package binchunk

const (
	LUA_VERSION_MAJOR   = "5"
	LUA_VERSION_MINOR   = "3"
	LUA_VERSION_NUM     = 503
	LUA_VERSION_RELEASE = "6"

	LUA_VERSION   = "Lua " + LUA_VERSION_MAJOR + "." + LUA_VERSION_MINOR
	LUA_RELEASE   = LUA_VERSION + "." + LUA_VERSION_RELEASE
	LUA_COPYRIGHT = LUA_RELEASE + "  Copyright (C) 1994-2020 Lua.org, PUC-Rio"
	LUA_AUTHORS   = "R. Ierusalimschy, L. H. de Figueiredo, W. Celes"

	LUA_SIGNATURE = "\x1bLua"
)

const (
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
)

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)
