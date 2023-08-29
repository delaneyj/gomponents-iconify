package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Traccarclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.879 34.962a6.958 6.958 0 0 1-9.84-9.84l4.92 4.92Z"/><circle cx="23.923" cy="24.077" r="1.193" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.303 17.539a6.958 6.958 0 0 1 4.158 4.158M28.682 11A13.916 13.916 0 0 1 37 19.318"/>`),
		g.Group(children),
	)
}