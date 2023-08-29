package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 24v-4a2.002 2.002 0 0 1 2-2h15a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2H6a2.002 2.002 0 0 1-2-2zm2 0h15v-4l-15-.001zm6-12V8a2.002 2.002 0 0 1 2-2h7a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2h-7a2.002 2.002 0 0 1-2-2zm2 0h7V8l-7-.001zm14 18h-2V2h2z"/>`),
		g.Group(children),
	)
}