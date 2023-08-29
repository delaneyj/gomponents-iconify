package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SectionRemoveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 .5H1.5a1 1 0 0 0-1 1V4M6 .5h3m2 0h2.5a1 1 0 0 1 1 1V4M.5 6v3m14-3v3m-14 2v2.5a1 1 0 0 0 1 1H4M14.5 11v2.5a1 1 0 0 1-1 1H11m-7-7h7m-5 7h3"/>`),
		g.Group(children),
	)
}