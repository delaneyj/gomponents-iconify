package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 2h-2V0h-2v2h-2v2h2v2h2V4h2z"/><path fill="currentColor" d="M13.498 6.969c.288.32.55.665.782 1.031a7.594 7.594 0 0 1-2.335 2.348a7.326 7.326 0 0 1-7.889 0A7.626 7.626 0 0 1 1.721 8a7.594 7.594 0 0 1 2.52-2.462A4 4 0 1 0 12 6.907v-.032a4.002 4.002 0 0 1-2.999-3.817A8.94 8.94 0 0 0 8 3.001c-3.489 0-6.514 2.032-8 5c1.486 2.968 4.511 5 8 5s6.514-2.032 8-5a9.217 9.217 0 0 0-.979-1.548a3.973 3.973 0 0 1-1.523.517zM6.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 6.5 5z"/>`),
		g.Group(children),
	)
}