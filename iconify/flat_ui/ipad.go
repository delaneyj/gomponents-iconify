package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#2C3E50" fill-rule="evenodd" d="M5 0h68a5 5 0 0 1 5 5v90a5 5 0 0 1-5 5H5a5 5 0 0 1-5-5V5a5 5 0 0 1 5-5z" clip-rule="evenodd"/><path fill="#35495E" fill-rule="evenodd" d="M4.991 11v71h68V11h-68zm34 76a4 4 0 1 0 0 8a4 4 0 0 0 0-8z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}