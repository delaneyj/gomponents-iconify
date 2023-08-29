package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TtbTouch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.993 32.942c3.8 0 6.507-2.95 6.507-6.618s-2.708-6.496-6.507-6.496s-6.485 2.829-6.485 6.496s2.685 6.618 6.485 6.618Zm-6.485-17.884v4.703m-13.015 0v-4.703m13.015 4.703v6.867a6.058 6.058 0 0 1-6.045 6.071h-.899a6.058 6.058 0 0 1-6.071-6.045v-6.893h5.242m-18.235 0v-4.703m12.993 11.57a6.04 6.04 0 0 1-6.009 6.07h-.913a6.058 6.058 0 0 1-6.07-6.045V19.76h5.24"/>`),
		g.Group(children),
	)
}