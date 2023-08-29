package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonePaused(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 0v149h-43V0h43zm64 267q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12zM341 0h43v149h-43V0z"/>`),
		g.Group(children),
	)
}