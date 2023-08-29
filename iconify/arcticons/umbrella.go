package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.806 40.837a3.097 3.097 0 0 0 6.194 0V23.566m-19.5-.001a19.5 19.5 0 0 1 39 0a4.875 4.875 0 1 0-9.75 0a4.875 4.875 0 1 0-9.75 0a4.875 4.875 0 1 0-9.75 0a4.875 4.875 0 1 0-9.75 0Z"/>`),
		g.Group(children),
	)
}