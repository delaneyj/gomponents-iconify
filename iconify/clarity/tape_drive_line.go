package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapeDriveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityTapeDriveLine0" fill="currentColor"><path d="M32 6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2ZM4 28V8h28v20Z"/><path d="M13.33 13.35a4.52 4.52 0 1 0 4.53 4.52a4.53 4.53 0 0 0-4.53-4.52Zm0 7.44a2.92 2.92 0 1 1 2.93-2.92a2.92 2.92 0 0 1-2.93 2.92Zm10.29-7.44a4.52 4.52 0 1 0 4.52 4.52a4.53 4.53 0 0 0-4.52-4.52Zm0 7.44a2.92 2.92 0 1 1 2.92-2.92a2.92 2.92 0 0 1-2.92 2.92Z"/><path d="M6 11v12.55h2V12h21.34v-2H7a1 1 0 0 0-1 1Z"/></g>`),
		g.Group(children),
	)
}