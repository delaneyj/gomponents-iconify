package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func O(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 96a160 160 0 1 0 0 320a160 160 0 1 0 0-320zm224 160a224 224 0 1 1-448 0a224 224 0 1 1 448 0z"/>`),
		g.Group(children),
	)
}