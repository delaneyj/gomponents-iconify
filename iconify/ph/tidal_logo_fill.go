package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TidalLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m253.66 101.66l-36 36a8 8 0 0 1-11.32 0l-36-36l-.34-.38l-.34.38L135.31 136l34.35 34.34a8 8 0 0 1 0 11.32l-36 36a8 8 0 0 1-11.32 0l-36-36a8 8 0 0 1 0-11.32L120.69 136l-34.35-34.34l-.34-.38l-.34.38l-36 36a8 8 0 0 1-11.32 0l-36-36a8 8 0 0 1 0-11.32l36-36a8 8 0 0 1 11.32 0l36 36l.34.38l.34-.38l36-36a8 8 0 0 1 11.32 0l36 36l.34.38l.34-.38l36-36a8 8 0 0 1 11.32 0l36 36a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}