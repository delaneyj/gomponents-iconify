package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buckwheat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.366 5.5S-6.242 15.648 9.757 32.142C28.017 50.967 50.24 45.305 38.366 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.371 5.57c-4.418 18.053-6.16 30.306-19.218 33.913"/>`),
		g.Group(children),
	)
}