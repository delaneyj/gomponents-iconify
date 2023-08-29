package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#0052b4" d="M0 0h512v256H0z"/><path fill="#d80027" d="M0 256h512v256H0z"/><path fill="#eee" d="m256 0l48 208l208 48l-208 48l-48 208l-48-208L0 256l208-48Z"/></g>`),
		g.Group(children),
	)
}