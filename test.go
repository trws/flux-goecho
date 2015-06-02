package main

// #cgo pkg-config: flux-core json-c
// #cgo LDFLAGS: -lmrpc
// #include <stdlib.h>
// #include <errno.h>
// #include <flux/core.h>
// #include <flux/mrpc.h>
// #include <json.h>
// extern const char * mod_name;
// int mecho_mrpc_cb_gateway (flux_t h, int typemask, zmsg_t **zmsg, void *arg);
// void flux_log_wrapper(flux_t h, int level, const char * str);
// void print_crap();
import "C"
import "fmt"
import "unsafe"
// import "os"

func main() {
  fmt.Println("in module's random main?")
}

func flux_log(h C.flux_t, level C.int, args ...interface{}) {
  str := C.CString(fmt.Sprintf(args[0].(string), args[1:]...))
  defer C.free(unsafe.Pointer(str))

  C.flux_log_wrapper(h, level, str)
}

//ensure a C-callable prototype exists
//export mecho_mrpc_cb
func mecho_mrpc_cb (h C.flux_t, typemask C.int, zmsg **C.zmsg_t , arg unsafe.Pointer) C.int {
  //clean up input arg
  defer C.zmsg_destroy (zmsg)

  var json_str *C.char = nil
  var inarg *C.json_object = nil
  ret, err := C.flux_event_decode (*zmsg, nil, &json_str)
  if ret < 0 {
    flux_log (h, C.LOG_ERR, "flux_event_decode: %v", err)
    return -1
  }
  request, err := C.json_tokener_parse (json_str)
  if request == nil {
    flux_log (h, C.LOG_ERR, "tokener_parse: %v", err)
    return -1
  }
  defer C.json_object_put (request)

  f, err := C.flux_mrpc_create_fromevent (h, request)
  if f == nil {
    flux_log (h, C.LOG_ERR, "flux_mrpc_create_fromevent: %v", err)
    return -1
  }
  defer C.flux_mrpc_destroy (f)

  if C.flux_mrpc_get_inarg (f, &inarg) < 0 {
    flux_log (h, C.LOG_ERR, "flux_mrpc_get_inarg: %v", err);
    return -1
  }
  defer C.json_object_put (inarg)

  _, err = C.flux_mrpc_put_outarg (f, inarg)
  if C.flux_mrpc_respond (f) < 0 {
    flux_log (h, C.LOG_ERR, "flux_mrpc_respond: %v", err);
    return -1
  }
  return 0
}

//export mod_main
func mod_main(h C.flux_t , args *C.zhash_t ) C.int {
  modname := C.CString("mrpc.mecho")
  defer C.free(unsafe.Pointer(modname))
  if C.flux_event_subscribe (h, modname) < 0 {
    // no variadic function support, don't feel like interposing this
    fmt.Println ("%s: flux_event_subscribe", "mod_main")
    return -1
  }
  ret, err := C.flux_msghandler_add (h, C.FLUX_MSGTYPE_EVENT, modname, C.FluxMsgHandler(unsafe.Pointer(C.mecho_mrpc_cb_gateway)), nil)
  if ret < 0 {
    fmt.Println ("flux_msghandler_add: %d", err)
    return -1
  }
  ret, err = C.flux_reactor_start (h)
  if ret < 0 {
    fmt.Println ("flux_reactor_start: %d", err)
    return -1
  }
  return 0
}
