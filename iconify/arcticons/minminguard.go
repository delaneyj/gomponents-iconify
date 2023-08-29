package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minminguard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.07 4.5H15.93L4.5 15.93v16.14L15.93 43.5h16.14L43.5 32.07V15.93Zm7.29 17.9a3.24 3.24 0 0 0-3.42-3.23a3.36 3.36 0 0 0-3.06 3.42v3a3.24 3.24 0 0 0 3.24 3.24h0a3.24 3.24 0 0 0 3.24-3.24h-3.24M8.64 28.83v-9.67l4.83 9.68l4.82-9.66v9.66m2.47 0v-9.68l4.82 9.68l4.82-9.66v9.66"/>`),
		g.Group(children),
	)
}