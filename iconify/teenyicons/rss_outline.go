package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 13.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm14 1.5C14.5 6.992 8.008.5 0 .5m0 6A8.5 8.5 0 0 1 8.5 15"/>`),
		g.Group(children),
	)
}