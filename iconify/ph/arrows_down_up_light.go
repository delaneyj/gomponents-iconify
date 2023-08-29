package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsDownUpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M116.24 171.76a6 6 0 0 1 0 8.48l-32 32a6 6 0 0 1-8.48 0l-32-32a6 6 0 0 1 8.48-8.48L74 193.51V48a6 6 0 0 1 12 0v145.51l21.76-21.75a6 6 0 0 1 8.48 0Zm96-96l-32-32a6 6 0 0 0-8.48 0l-32 32a6 6 0 0 0 8.48 8.48L170 62.49V208a6 6 0 0 0 12 0V62.49l21.76 21.75a6 6 0 0 0 8.48-8.48Z"/>`),
		g.Group(children),
	)
}