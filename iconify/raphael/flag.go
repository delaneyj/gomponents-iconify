package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.5 3v10c8 0 8 4 16 4V7c-8 0-8-4-16-4zm-3 26h2V3h-2v26z"/>`),
		g.Group(children),
	)
}