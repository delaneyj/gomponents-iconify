package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperscriptOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m12.5 3.5l-.354-.354A.5.5 0 0 0 12.5 4v-.5Zm1.793-1.793l-.354-.353l.354.353ZM15 3h-2.5v1H15V3Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM13.793 0H13.5v1h.293V0ZM13.5 0A1.5 1.5 0 0 0 12 1.5h1a.5.5 0 0 1 .5-.5V0ZM15 1.207C15 .54 14.46 0 13.793 0v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147l.706.707ZM1.1 1.8l9 12l.8-.6l-9-12l-.8.6Zm9-.6l-9 12l.8.6l9-12l-.8-.6Z"/>`),
		g.Group(children),
	)
}