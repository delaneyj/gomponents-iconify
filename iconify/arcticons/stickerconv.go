package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stickerconv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.611 45.44q-.798.06-1.611.06A21.5 21.5 0 1 1 45.5 24q0 .766-.053 1.518"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.564 45.326a21.504 21.504 0 0 1 19.868-19.77m-17.633 4.278l-4.058 4.058l-5.575-5.626l-5.575 5.626l-4.058-4.058l9.633-9.633Zm-7.598-11.668l4.058-4.058l5.575 5.626l5.575-5.626l4.058 4.058l-9.633 9.633Zm5.41 27.274l19.821-19.884"/>`),
		g.Group(children),
	)
}