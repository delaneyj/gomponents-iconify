package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftUppercaseSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.817 1.571a1.511 1.511 0 0 1 2.367 0l4.598 5.802c.52.657.05 1.621-.789 1.621h-1.978V11c0 .553-.45 1.001-1.005 1.001H5.99A1.003 1.003 0 0 1 4.985 11V8.994H3.007c-.84 0-1.31-.965-.789-1.62l4.6-5.803ZM5.5 13a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Z"/>`),
		g.Group(children),
	)
}