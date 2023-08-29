package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageNumberOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-4 9l-.354-.354A.5.5 0 0 0 9.5 13v-.5Zm3 1.5h-10v1h10v-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2.5 0A1.5 1.5 0 0 0 1 1.5h1a.5.5 0 0 1 .5-.5V0ZM12 12H9.5v1H12v-1Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM10.793 9H10.5v1h.293V9ZM10.5 9A1.5 1.5 0 0 0 9 10.5h1a.5.5 0 0 1 .5-.5V9Zm1.5 1.207C12 9.54 11.46 9 10.793 9v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147l.706.707ZM13 3.5v10h1v-10h-1Zm-11 10v-12H1v12h1Z"/>`),
		g.Group(children),
	)
}