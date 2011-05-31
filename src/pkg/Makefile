# Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all: install

DIRS=\
     solox/core\
     solox/build\
     lang\

NOTEST=\

NOBENCH=\

TEST=\
     	$(filter-out $(NOTEST),$(DIRS))

BENCH=\
      	$(filter-out $(NOBENCH),$(TEST))


clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))
nuke.dirs: $(addsuffix .nuke, $(DIRS)) 
test.dirs: $(addsuffix .test, $(TEST))
bench.dirs: $(addsuffix .bench, $(BENCH))

%.clean:
		+cd $* && gomake clean

%.install:
		+cd $* && gomake install

%.nuke:
		+cd $* && gomake nuke

%.test:
		+cd $* && gomake test

%.bench:
		+cd $* && gomake bench


clean: clean.dirs

install: install.dirs

test:	test.dirs

#bench: bench.dirs path/to/file.bench

nuke: nuke.dirs
#	rm -rf $(GOROOT)/pkg/*
