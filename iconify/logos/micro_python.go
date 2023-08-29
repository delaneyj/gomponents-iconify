package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroPython(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M0 0h122.76v199.86h10.48V0H256v256h-56.14V56.14h-10.48V256H66.62V56.14l-10.48.375V256H0V0Zm237.287 208.094h-14.971v25.45h14.97v-25.45Z"/>`),
		g.Group(children),
	)
}