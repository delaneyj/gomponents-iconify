package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unitconverterultimate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.66 5.69v5.24H31v8.66H13.69l-.18 4.95L4.5 15l9.16-9.34Zm20.41 17.76l9.43 9.43l-9.43 9.43v-5.1H16.75v-8.66h17.32v-5.1Z"/>`),
		g.Group(children),
	)
}