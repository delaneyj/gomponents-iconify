package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M64 3q88 0 150.5 62.5T277 216t-62.5 150.5T64 429q-33 0-64-9q66-21 107.5-77T149 216T107.5 89T0 12q31-9 64-9z"/>`),
		g.Group(children),
	)
}