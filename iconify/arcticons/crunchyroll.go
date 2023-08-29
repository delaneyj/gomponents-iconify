package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crunchyroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 46a22 22 0 1 1 22-22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.218 26.001a15.714 15.714 0 1 1-8.873-12.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.117 25.912a7.571 7.571 0 1 1-8.713-12.148"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.876 45.093a20.285 20.285 0 1 1 26.818-28.408"/>`),
		g.Group(children),
	)
}