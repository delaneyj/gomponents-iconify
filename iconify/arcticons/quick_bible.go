package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickBible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 42.138h39M24 36.994V5.862m-9.533 8.345h19.066M20.98 40.438L4.953 39.65l.615-4.542Zm-2.565-2.991l-9.86-5.598m18.465 8.589l16.028-.788l-.615-4.542Zm2.565-2.991l9.86-5.598"/>`),
		g.Group(children),
	)
}