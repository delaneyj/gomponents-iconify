package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1186 925l257-250l-356-52l-66-10l-30-60l-159-322v963l59 31l318 168l-60-355l-12-66zm452-262l-363 354l86 500q5 33-6 51.5t-34 18.5q-17 0-40-12l-449-236l-449 236q-23 12-40 12q-23 0-34-18.5t-6-51.5l86-500L25 663q-32-32-23-59.5T56 569l502-73L783 41q20-41 49-41q28 0 49 41l225 455l502 73q45 7 54 34.5t-24 59.5z"/>`),
		g.Group(children),
	)
}