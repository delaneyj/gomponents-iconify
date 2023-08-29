package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathOperationsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M110 72a6 6 0 0 1-6 6H40a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Zm-6 106H78v-26a6 6 0 0 0-12 0v26H40a6 6 0 0 0 0 12h26v26a6 6 0 0 0 12 0v-26h26a6 6 0 0 0 0-12Zm48-4h64a6 6 0 0 0 0-12h-64a6 6 0 0 0 0 12Zm64 20h-64a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm-60.24-93.76a6 6 0 0 0 8.48 0L184 80.49l19.76 19.75a6 6 0 0 0 8.48-8.48L192.49 72l19.75-19.76a6 6 0 0 0-8.48-8.48L184 63.51l-19.76-19.75a6 6 0 0 0-8.48 8.48L175.51 72l-19.75 19.76a6 6 0 0 0 0 8.48Z"/>`),
		g.Group(children),
	)
}