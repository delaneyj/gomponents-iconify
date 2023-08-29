package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#23475F" d="M91 62V5a5 5 0 0 0-5-5H14a5 5 0 0 0-5 5v57H0v3a4 4 0 0 0 4 4h92a4 4 0 0 0 4-4v-3h-9z"/><path fill="#3498DB" d="M14 5h72v51H14V5z"/><path fill="#9ACCED" d="M58 19H43v4h15v-4zm-25 8v2h35v-2H33zm0 6h35v-2H33v2zm0 4h35v-2H33v2zm5 5.001h25V40H38v2.001z"/><path fill="#1C394C" d="M9 60h82v2H9v-2z"/><path fill="#BCC4C8" d="M74.5 64a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 74.5 64zm5 0a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 79.5 64zm5 0a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 84.5 64z"/>`),
		g.Group(children),
	)
}