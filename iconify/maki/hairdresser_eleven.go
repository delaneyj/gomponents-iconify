package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HairdresserEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M1.5 2A1.5 1.5 0 0 0 0 3.5v1c0 .5.5.5.5.5h1C3 5 4 5.5 4 5.5S3 6 1.5 6h-1S0 6 0 6.5v1a1.5 1.5 0 0 0 3 0v-.615c.808-.158 1.587-.453 2.225-.742L8.5 8C10 8.75 11 8 11 8L6.5 5.5L11 3s-1-.75-2.5 0L5.225 4.857C4.587 4.568 3.808 4.273 3 4.115v-.611V3.5A1.5 1.5 0 0 0 1.5 2zm0 1a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zm3.732 2.25h.018a.25.25 0 1 1-.018 0zM1.5 7a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}