package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GavelFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M52.69 99.31a16 16 0 0 1 0-22.63l64-64a16 16 0 0 1 22.63 22.63l-64 64a16 16 0 0 1-22.63 0Zm190.63 17.37a16 16 0 0 0-22.63 0l-64 64a16 16 0 0 0 0 22.62a16 16 0 0 0 22.63 0l64-64a16 16 0 0 0 0-22.62Zm-37.66-18.34l-48-48a8 8 0 0 0-11.31 0l-56 56a8 8 0 0 0 0 11.31L103 130.34l-74.31 74.35a16 16 0 0 0 22.62 22.62L125.66 153l12.69 12.69a8 8 0 0 0 11.31 0l56-56a8 8 0 0 0 0-11.35Z"/>`),
		g.Group(children),
	)
}