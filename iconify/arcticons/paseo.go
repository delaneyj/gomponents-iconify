package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paseo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.681 5.918l3.725 20.016l-.627.054c-.266 3.105-1.962 5.522-5.58 6.785s-8.381 3.988-7.891 7.376c1.317 4.03 4.468 4.543 15.162.895s14.603-4.154 14.233-5.81c-.491-2.2-3.604-9.456-5.001-9.06a6.83 6.83 0 0 0-.734.216L28.766 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.408 25.945c1.5.71 8.086.644 9.612.459"/>`),
		g.Group(children),
	)
}