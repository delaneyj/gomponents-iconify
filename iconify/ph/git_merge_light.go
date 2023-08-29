package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMergeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 114a30.05 30.05 0 0 0-29.4 24H152a42.2 42.2 0 0 1-33.6-16.8L90.56 84.08A30 30 0 1 0 74 85.4v85.2a30 30 0 1 0 12 0V98l22.8 30.4A54.26 54.26 0 0 0 152 150h26.6a30 30 0 1 0 29.4-36ZM62 56a18 18 0 1 1 18 18a18 18 0 0 1-18-18Zm36 144a18 18 0 1 1-18-18a18 18 0 0 1 18 18Zm110-38a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}