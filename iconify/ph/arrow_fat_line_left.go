package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 72h-56V32a8 8 0 0 0-13.66-5.66l-96 96a8 8 0 0 0 0 11.32l96 96A8 8 0 0 0 128 224v-40h56a8 8 0 0 0 8-8V80a8 8 0 0 0-8-8Zm-8 96h-56a8 8 0 0 0-8 8v28.69L35.31 128L112 51.31V80a8 8 0 0 0 8 8h56Zm48-88v96a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}