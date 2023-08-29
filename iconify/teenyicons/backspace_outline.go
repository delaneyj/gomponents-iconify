package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m4.5 12.5l-.39.312A.5.5 0 0 0 4.5 13v-.5Zm-4-5l-.39-.312a.5.5 0 0 0 0 .624L.5 7.5Zm4-5V2a.5.5 0 0 0-.39.188l.39.312Zm9.5 1v8h1v-8h-1Zm-.5 8.5h-9v1h9v-1Zm-8.61.188l-4-5l-.78.624l4 5l.78-.624Zm-4-4.376l4-5l-.78-.624l-4 5l.78.624ZM4.5 3h9V2h-9v1Zm9.5 8.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1Zm1-8A1.5 1.5 0 0 0 13.5 2v1a.5.5 0 0 1 .5.5h1ZM6.146 5.854l4 4l.708-.708l-4-4l-.708.708Zm4-.708l-4 4l.708.708l4-4l-.708-.708Z"/>`),
		g.Group(children),
	)
}