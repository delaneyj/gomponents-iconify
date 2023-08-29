package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 7H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4zm0 7H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}