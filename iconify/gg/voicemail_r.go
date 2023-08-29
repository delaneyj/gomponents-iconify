package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.5 12c0 .35-.06.687-.17 1h2.34a3 3 0 1 1 2.83 2h-8a3 3 0 1 1 3-3Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm8 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M1.5 8a3 3 0 0 1 3-3h15a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-15a3 3 0 0 1-3-3V8Zm3-1h15a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1Z"/></g>`),
		g.Group(children),
	)
}