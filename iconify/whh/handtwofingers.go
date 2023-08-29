package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handtwofingers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024H128q-30 0-61-51T18 862T0 768V512q0-26 19-45t45-19q19 0 36 12L64 70q-2-28 15-48.5T122 0t46 17.5T189 63l35 321q5 11 14.5 21.5T256 416l35-353q1-27 20.5-45.5T357 0t43.5 21.5T416 70q-32 186-32 378q0-26 18.5-45t45-19t45.5 19t19 45v64q0-26 18.5-45t45-19t45.5 19t19 45v224q0 63-13 115.5t-32 83t-38 52t-32 29.5z"/>`),
		g.Group(children),
	)
}