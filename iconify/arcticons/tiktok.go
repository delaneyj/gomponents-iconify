package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiktok(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.73 22.4A10.55 10.55 0 1 0 29.27 33V4.5a10.55 10.55 0 0 0 10.55 10.55"/>`),
		g.Group(children),
	)
}