package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 12h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1zm0 4h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1zm3-11h8a.5.5 0 0 0 0-1h-8a.5.5 0 0 0 0 1zm-3 15h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1zm0-12h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1zm7 12h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1zm4-8h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1zm-11-8h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1zm11 4h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1zm0 8h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}