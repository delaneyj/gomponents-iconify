package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Temu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.761 40.397a7.739 7.739 0 0 0 7.726-8.187l-.715-12.325c-.4-6.895-6.107-12.282-13.014-12.282H18.242c-6.907 0-12.614 5.387-13.014 12.282L4.513 32.21a7.739 7.739 0 0 0 7.726 8.187h23.522Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.811 16.91a7.188 7.188 0 1 0 14.378 0"/>`),
		g.Group(children),
	)
}