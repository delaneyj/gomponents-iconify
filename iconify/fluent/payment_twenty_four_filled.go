package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 5A3.25 3.25 0 0 0 2 8.25V9.5h20V8.25A3.25 3.25 0 0 0 18.75 5H5.25ZM22 11H2v4.75A3.25 3.25 0 0 0 5.25 19h13.5A3.25 3.25 0 0 0 22 15.75V11Zm-6.25 3.5h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}