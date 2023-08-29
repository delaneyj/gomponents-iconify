package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.25 6H7.444l-.125.996h.094c.324 0 .745.003.906.026a2.25 2.25 0 1 1-2.248 3.377a.5.5 0 1 1 .86-.511a1.25 1.25 0 1 0 1.25-1.875l-.007-.001l-.018-.002a25.465 25.465 0 0 0-.743-.014c-.171 0-.335 0-.456.002l-.145.001h-.053a.502.502 0 0 1-.503-.561l.25-2A.5.5 0 0 1 7.001 5h2.25a.5.5 0 0 1 0 1ZM8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2ZM3 8a5 5 0 1 1 10 0A5 5 0 0 1 3 8Z"/>`),
		g.Group(children),
	)
}