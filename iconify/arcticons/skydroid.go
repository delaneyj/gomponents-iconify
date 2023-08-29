package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skydroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.47 26.5H42.5v16H26.47zm-20.97 0h16v16h0a16 16 0 0 1-16-16v0h0Zm0-21h16v0a16 16 0 0 1-16 16h0v-16h0Zm20.97 0h0a16 16 0 0 1 16 16v0h0h-16h0v-16h0Z"/>`),
		g.Group(children),
	)
}