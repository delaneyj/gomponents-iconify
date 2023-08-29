package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11 12c0 .35-.06.687-.17 1h2.34A3 3 0 1 1 16 15H8a3 3 0 1 1 3-3Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm8 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/></g>`),
		g.Group(children),
	)
}