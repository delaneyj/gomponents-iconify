package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17.845 6.155s-.433 2.245-2.94 4.751c-2.505 2.506-4.75 2.938-4.75 2.938m10.252 2.563a5.437 5.437 0 0 1-7.688 0l-5.127-5.126a5.437 5.437 0 1 1 7.69-7.689l5.125 5.126a5.437 5.437 0 0 1 0 7.69Z"/><path stroke-linecap="round" d="M14.5 6.5L13 5" opacity=".5"/><path d="M6.73 10.135a6 6 0 1 0 7.143 7.098" opacity=".5"/></g>`),
		g.Group(children),
	)
}