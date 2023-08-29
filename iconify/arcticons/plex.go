package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.22 4.5h12.13s9.31 15.12 11 18a3.19 3.19 0 0 1 0 3l-11 18H12.43S22 28.06 23.05 26.38a3.73 3.73 0 0 0 0-4.55C21.62 19.53 12.22 4.5 12.22 4.5Z"/>`),
		g.Group(children),
	)
}