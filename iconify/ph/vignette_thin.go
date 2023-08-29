package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VignetteThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44H40a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm4 156a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4V56a4 4 0 0 1 4-4h176a4 4 0 0 1 4 4ZM128 76c-37.5 0-68 23.33-68 52s30.5 52 68 52s68-23.33 68-52s-30.5-52-68-52Zm0 96c-33.08 0-60-19.74-60-44s26.92-44 60-44s60 19.74 60 44s-26.92 44-60 44Z"/>`),
		g.Group(children),
	)
}