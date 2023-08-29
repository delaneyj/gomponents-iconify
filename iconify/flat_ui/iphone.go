package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#2C3E50" fill-rule="evenodd" d="M5 0h50a5 5 0 0 1 5 5v90a5 5 0 0 1-5 5H5a5 5 0 0 1-5-5V5a5 5 0 0 1 5-5z" clip-rule="evenodd"/><path fill="#35495E" fill-rule="evenodd" d="M5 11h50v73H5V11zm25 77a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}