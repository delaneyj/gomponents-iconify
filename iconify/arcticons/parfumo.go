package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parfumo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.287 2.823c-3.647 13.312-4.241 23.245-2.015 41.902"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.522 11.476c6.071 4 8.755 6.039 10.566 10.102m26.195-7.094c-4.996 8.992-8.281 22.26-9.133 28.473"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.026 6.896c2.485 5.569 4.249 11.614 4.249 11.614"/>`),
		g.Group(children),
	)
}