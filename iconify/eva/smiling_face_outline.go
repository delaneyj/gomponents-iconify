package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFaceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSmilingFaceOutline0"><g id="evaSmilingFaceOutline1"><path id="evaSmilingFaceOutline2" fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm5 9a5 5 0 0 1-10 0Z"/></g></g>`),
		g.Group(children),
	)
}