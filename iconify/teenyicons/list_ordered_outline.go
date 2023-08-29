package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOrderedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 13.5l-.354-.354A.5.5 0 0 0 .5 14v-.5Zm1-11H2V2h-.5v.5ZM5 8h10V7H5v1Zm0-4h10V3H5v1Zm0 8h10v-1H5v1Zm-2 1H.5v1H3v-1Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM1.793 10H1.5v1h.293v-1ZM1.5 10A1.5 1.5 0 0 0 0 11.5h1a.5.5 0 0 1 .5-.5v-1ZM3 11.207C3 10.54 2.46 10 1.793 10v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854H2a.207.207 0 0 1-.06.147l.706.707ZM0 6h3V5H0v1Zm2-.5v-3H1v3h1ZM1.5 2H0v1h1.5V2Z"/>`),
		g.Group(children),
	)
}