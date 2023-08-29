package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9.874 5.008c2.728-1.68 6.604-1.014 8.25.197c-2.955.84-5.11 3.267-5.242 6.415c-.18 4.28 3.006 6.588 5.24 7.152c-1.964 1.343-4.36 1.293-5.235 1.172c-3.568-.492-6.902-3.433-7.007-7.711c-.106-4.278 2.573-6.35 3.994-7.225z"/>`),
		g.Group(children),
	)
}