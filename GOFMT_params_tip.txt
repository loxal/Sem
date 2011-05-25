# Copyright 2010 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

GOFMT=gofmt -spaces=true -tabindent=true -tabwidth=18

TARG=solox/domain_lang_extractor
GOFILES=\
	domain_lang_extractor.go\

include $(GOROOT)/src/Make.pkg

format:
		${GOFMT} -w domain_lang_extractor.go
		${GOFMT} -w main.go
