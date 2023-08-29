package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDEllipseThreePts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M5 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M5 22h8m-8 0V2"/><path fill="currentColor" d="M5 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path stroke-dasharray="3 3" d="M8 4.193C9.37 2.821 11.108 2 13 2c4.418 0 8 4.477 8 10c0 3.271-1.256 6.175-3.2 8"/><path d="M8.2 20A9.098 9.098 0 0 1 7 18.615"/><path fill="currentColor" d="M13 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`),
		g.Group(children),
	)
}