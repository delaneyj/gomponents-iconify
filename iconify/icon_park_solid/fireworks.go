package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fireworks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFireworks0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m6 42l8.674-24.736L31 34.038L6 42Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m23 19l5-5c2.667-2.667 3-5 1-7m0 18l5-5c3.333-3.333 6.667-3.333 10 0"/><path fill="#fff" d="M20 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm22-1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 23a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-3 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFireworks0)"/>`),
		g.Group(children),
	)
}