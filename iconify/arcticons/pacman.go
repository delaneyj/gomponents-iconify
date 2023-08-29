package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pacman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24l19.93-8.052h0a21.494 21.494 0 1 0-.144 16.45Z"/><circle cx="37.247" cy="24" r="1.751" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="43.749" cy="24" r="1.751" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}