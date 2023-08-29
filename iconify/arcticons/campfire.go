package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campfire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.854 44.278c2.531-9.694 12.415-8.001 14.226.023"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.986 37.181C36.876 28.991 30.87 21.325 24 13.9c-6.639 6.685-13.009 15.17-17.122 23.105"/>`),
		g.Group(children),
	)
}