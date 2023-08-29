package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckFatLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m241.87 69.66l-24-23.56a14 14 0 0 0-19.77 0l-92.72 91.34a2 2 0 0 1-2.83 0l-.09-.08l-36.61-35.31a14 14 0 0 0-19.75 0l-24 24a14 14 0 0 0 0 19.79l71.62 72a14 14 0 0 0 19.8 0L241.91 89.5a14 14 0 0 0-.04-19.84ZM233.42 81L105 209.41a2 2 0 0 1-2.81 0l-71.62-72a2 2 0 0 1 0-2.82l24-24A2 2 0 0 1 56 110a2 2 0 0 1 1.41.58l.08.08l36.66 35.28a14 14 0 0 0 19.72 0l92.72-91.35a2 2 0 0 1 2.87 0l24 23.56a2 2 0 0 1-.04 2.85Z"/>`),
		g.Group(children),
	)
}