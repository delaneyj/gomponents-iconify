package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAmplify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m112.31 268l40.36-68.69l34.65 59l-67.54 115h135L289.31 432H16Zm58.57-99.76l33.27-56.67L392.44 432h-66.68ZM222.67 80h66.59L496 432h-66.68Z"/>`),
		g.Group(children),
	)
}