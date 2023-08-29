package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneEqualizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 9h4v11h-4zm-6-5h4v16h-4zm-6 8h4v8H4z"/>`),
		g.Group(children),
	)
}