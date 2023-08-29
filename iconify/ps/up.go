package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Up(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m44 273l156-139l156 137q4 4 15 4q10 0 17-6q13-15-2-30L200 79L14 241q-14 16-2 30q14 13 32 2z"/>`),
		g.Group(children),
	)
}