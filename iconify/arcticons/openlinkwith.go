package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openlinkwith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.263 34.903a4.298 4.298 0 1 1-4.298 4.299a4.303 4.303 0 0 1 4.298-4.299ZM15.737 4.5a4.298 4.298 0 1 1-4.298 4.298A4.298 4.298 0 0 1 15.737 4.5Zm.658 30.403a4.298 4.298 0 1 1-4.29 4.299a4.298 4.298 0 0 1 4.29-4.299ZM32.263 4.5a4.298 4.298 0 1 1-4.298 4.298A4.298 4.298 0 0 1 32.263 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.507 34.903c-.658-13.886-16.103-7.92-16.77-21.806m16.526 0c0 3.965-1.102 7.323-3.307 9.083M16.395 34.903c0-3.97.778-7.16 2.982-8.92"/>`),
		g.Group(children),
	)
}