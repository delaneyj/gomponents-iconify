package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeBackMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M7 1024h1656l-674 674l144 145l922-921L1133 0L989 145l674 674H7v205z"/>`),
		g.Group(children),
	)
}