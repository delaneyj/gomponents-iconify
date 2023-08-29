package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TidalLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m250.83 93.17l-40-40a4 4 0 0 0-5.66 0L168 90.34l-37.17-37.17a4 4 0 0 0-5.66 0L88 90.34L50.83 53.17a4 4 0 0 0-5.66 0l-40 40a4 4 0 0 0 0 5.66l40 40a4 4 0 0 0 5.66 0L88 101.66L122.34 136l-37.17 37.17a4 4 0 0 0 0 5.66l40 40a4 4 0 0 0 5.66 0l40-40a4 4 0 0 0 0-5.66L133.66 136L168 101.66l37.17 37.17a4 4 0 0 0 5.66 0l40-40a4 4 0 0 0 0-5.66ZM48 130.34L13.66 96L48 61.66L82.34 96Zm80 80L93.66 176L128 141.66L162.34 176Zm0-80L93.66 96L128 61.66L162.34 96Zm80 0L173.66 96L208 61.66L242.34 96Z"/>`),
		g.Group(children),
	)
}