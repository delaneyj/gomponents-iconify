package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.002 4.553a50.577 50.577 0 0 1 8.099.04l1.623.138a2.666 2.666 0 0 1 2.409 2.252l.102.669a20.665 20.665 0 0 1-.096 6.835a2.343 2.343 0 0 1-2.305 1.923H8.858c-.207 0-.41.051-.592.149l-3.911 2.102A.75.75 0 0 1 3.25 18V9.483a4.93 4.93 0 0 1 4.559-4.915l.193-.015ZM8 9.5A1.25 1.25 0 1 0 8 12a1.25 1.25 0 0 0 0-2.5Zm4 0a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5Zm2.75 1.25a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}