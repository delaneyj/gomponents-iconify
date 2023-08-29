package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourlyreminder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.277 32.7V21.49A13.272 13.272 0 0 0 27.08 8.588v-.985a3.102 3.102 0 0 0-6.204 0V8.6a13.271 13.271 0 0 0-10.153 12.89V32.7l-4.206 4.206v1.943h34.966v-1.943Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.326 38.849a4.652 4.652 0 1 0 9.303 0"/>`),
		g.Group(children),
	)
}