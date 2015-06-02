package main

// #cgo pkg-config: flux-core
/*
//gateway function to pass as callback
#include <flux/core.h>

int mecho_mrpc_cb_gateway (flux_t h, int typemask, zmsg_t **zmsg, void *arg){
  return mecho_mrpc_cb(h, typemask, zmsg, arg);
}

void flux_log_wrapper(flux_t h, int level, const char * str){
  flux_log(h, level, str);
}
*/
import "C"
