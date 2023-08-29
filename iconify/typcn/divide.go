package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="6" r="2.25" fill="currentColor"/><circle cx="12" cy="18" r="2.25" fill="currentColor"/><path fill="currentColor" d="M18 10H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}