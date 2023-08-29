package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LouderSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.031 4.571V7.25C.031 8.118.697 9 1.521 9h6.447V3H1.521c-.824 0-1.49.703-1.49 1.571zm7.032 11.367H4.025L2.062 9.062H5.1l1.963 6.876zM13 0L9 2.76v5.981L13 12V0zm0 3.989V8c1.086 0 1.969-.897 1.969-2.006c0-1.107-.883-2.005-1.969-2.005z"/>`),
		g.Group(children),
	)
}