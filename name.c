#include <stdio.h>
#include <unistd.h>
#include <flux/core.h>

const char *mod_name = "gotest";

void print_crap(){
  printf("printing name from C: %s\n c ptr: %p\n", mod_name, mod_name);
}

__attribute__((constructor))
static void wait_out_init_on_open()
{
  //This should not be necessary, but dlclose during runtime init causes a
  //segfault
}
__attribute__((destructor))
static void wait_out_init_on_close()
{
}
