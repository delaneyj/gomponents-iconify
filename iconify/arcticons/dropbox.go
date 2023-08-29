package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="m33.81 19.36l9.69 6.17l-9.75 6.22L24 25.54l-9.75 6.21l-9.75-6.22l9.69-6.17l-9.69-6.17L14.25 7L24 13.19L33.75 7l9.75 6.21Zm-.13 0L24 13.2l-9.68 6.16L24 25.53ZM14.32 34.81l9.75-6.22l9.75 6.22L24.07 41Z"/>`),
		g.Group(children),
	)
}