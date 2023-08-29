package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M665 550L806 24h82L676 786L444 206L212 786L0 24h82l140 526L444 0z"/>`),
		g.Group(children),
	)
}