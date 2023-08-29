package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.5 1.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 6.5 1.5z"/><path fill="currentColor" d="m10 5l5.15-2.221l-.371-.929L8.5 4h-1L1.221 1.85l-.371.929L6 5v4l-2.051 6.634l.935.355L7.786 9.5h.429l2.902 6.489l.935-.355L10.001 9z"/>`),
		g.Group(children),
	)
}