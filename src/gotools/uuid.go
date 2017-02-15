package gotools

// #cgo LDFLAGS: -luuid
/*
#include <uuid/uuid.h>
#include <stdlib.h>
char * uuid_gen()
{
	char * uuid_str = (char *)malloc(37);
    // typedef unsigned char uuid_t[16];
    uuid_t uuid;

    // generate
    uuid_generate_time_safe(uuid);

    // unparse (to string)
    uuid_unparse_lower(uuid, uuid_str);
    //printf("generate uuid=%s\n", uuid_str);
    return (char *)uuid_str;
}
*/
import "C"

import "unsafe"

func UUID_Generate() string {
    cs := C.uuid_gen()
    defer C.free(unsafe.Pointer(cs))
    return C.GoString(cs)
}