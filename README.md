#Go mecho module, external build

If flux is installed in a system path, all you should need to do is make sure
you have go 1.5+ in your path with a proper `GOROOT`, run `make` and cp the
directory into the flux modules directory.  Currently the mrpc.h header is not
installed by default, so that needs to be added to
`<flux-prefix>/include/flux/` as well.

This module implements mecho in go, not terribly interesting perhaps, but
proves that it works.
