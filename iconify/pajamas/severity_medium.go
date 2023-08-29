package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeverityMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.04 10.454L.24 3.182c-.398-.603-.29-1.457.24-1.91C.688 1.097.94 1 1.2 1h9.6c.663 0 1.2.61 1.2 1.364c0 .295-.084.582-.24.818l-4.8 7.272c-.398.603-1.15.725-1.68.273a1.295 1.295 0 0 1-.24-.273Z"/>`),
		g.Group(children),
	)
}