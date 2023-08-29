package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 3A2.75 2.75 0 0 0 2 5.75V11a5 5 0 0 1 8 6v1h9.25A2.75 2.75 0 0 0 22 15.25v-9.5A2.75 2.75 0 0 0 19.25 3H4.75Zm2 4h10.5a.75.75 0 0 1 0 1.5H6.75a.75.75 0 0 1 0-1.5ZM12 12.75a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75ZM6 10a4 4 0 1 0 0 8.001A4 4 0 0 0 6 10Zm3 8.001c-.835.628-1.874 1-3 1a4.978 4.978 0 0 1-3-.998v3.246c0 .57.605.92 1.09.669l.09-.055L6 20.591l1.82 1.272a.75.75 0 0 0 1.172-.51L9 21.249L9.001 18Z"/>`),
		g.Group(children),
	)
}