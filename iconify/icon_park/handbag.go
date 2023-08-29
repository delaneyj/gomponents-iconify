package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handbag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 14C14 10.6863 16.6863 8 20 8H29C32.3137 8 35 10.6863 35 14V16H14V14Z"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 25L23.5149 29.8787C23.8334 29.9584 24.1666 29.9584 24.4851 29.8787L44 25V38C44 39.1046 43.1046 40 42 40H6C4.89543 40 4 39.1046 4 38V25Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 27V18C44 16.8954 43.1046 16 42 16H6C4.89543 16 4 16.8954 4 18V27"/><path fill="#000" d="M26.5 23C26.5 24.3807 25.3807 25.5 24 25.5C22.6193 25.5 21.5 24.3807 21.5 23C21.5 21.6193 22.6193 20.5 24 20.5C25.3807 20.5 26.5 21.6193 26.5 23Z"/></g>`),
		g.Group(children),
	)
}