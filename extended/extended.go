// Copyright 2020 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extended contains an extended set of terminal descriptions.
// Applications desiring to have a better chance of Just Working by
// default should include this package.  This will significantly increase
// the size of the program.
package extended

import (
	// The following imports just register themselves --
	// these are the terminal types we aggregate in this package.
	_ "github.com/ericwq/terminfo/a/aixterm"
	_ "github.com/ericwq/terminfo/a/alacritty"
	_ "github.com/ericwq/terminfo/a/ansi"
	_ "github.com/ericwq/terminfo/b/beterm"
	_ "github.com/ericwq/terminfo/c/cygwin"
	_ "github.com/ericwq/terminfo/d/dtterm"
	_ "github.com/ericwq/terminfo/e/emacs"
	// _ "github.com/ericwq/terminfo/f/foot"
	_ "github.com/ericwq/terminfo/g/gnome"
	_ "github.com/ericwq/terminfo/h/hpterm"
	_ "github.com/ericwq/terminfo/k/konsole"
	_ "github.com/ericwq/terminfo/k/kterm"
	_ "github.com/ericwq/terminfo/l/linux"
	_ "github.com/ericwq/terminfo/p/pcansi"
	_ "github.com/ericwq/terminfo/r/rxvt"
	_ "github.com/ericwq/terminfo/s/screen"
	_ "github.com/ericwq/terminfo/s/simpleterm"
	// _ "github.com/ericwq/terminfo/s/sun"
	_ "github.com/ericwq/terminfo/t/termite"
	_ "github.com/ericwq/terminfo/t/tmux"
	_ "github.com/ericwq/terminfo/v/vt100"
	_ "github.com/ericwq/terminfo/v/vt102"
	_ "github.com/ericwq/terminfo/v/vt220"
	_ "github.com/ericwq/terminfo/v/vt320"
	_ "github.com/ericwq/terminfo/v/vt400"
	_ "github.com/ericwq/terminfo/v/vt420"
	_ "github.com/ericwq/terminfo/v/vt52"
	_ "github.com/ericwq/terminfo/w/wy50"
	_ "github.com/ericwq/terminfo/w/wy60"
	_ "github.com/ericwq/terminfo/w/wy99_ansi"
	_ "github.com/ericwq/terminfo/x/xfce"
	_ "github.com/ericwq/terminfo/x/xterm"
	_ "github.com/ericwq/terminfo/x/xterm_kitty"
	// _ "github.com/ericwq/terminfo/x/xterm_termite"
)
