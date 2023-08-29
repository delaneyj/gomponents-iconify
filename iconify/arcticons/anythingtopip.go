package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anythingtopip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.601 30.483a82.503 82.503 0 0 1-11.238-.86a75.46 75.46 0 0 1 0-23.147a75.438 75.438 0 0 1 23.145 0a63.397 63.397 0 0 1 .861 10.775"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.637 41.524a75.438 75.438 0 0 1-23.145 0a75.456 75.456 0 0 1 0-23.146a75.434 75.434 0 0 1 23.145 0a75.46 75.46 0 0 1 0 23.146"/>`),
		g.Group(children),
	)
}