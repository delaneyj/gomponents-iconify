package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enrecipes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.752 40.645v-6.177a2.131 2.131 0 0 1 .638-1.526c.96-.943 2.891-2.848 4.988-4.971l-.009-.005a9.947 9.947 0 0 0-5.072-17.038a3.408 3.408 0 0 1-2.188-1.387a9.948 9.948 0 0 0-16.218 0a3.408 3.408 0 0 1-2.188 1.387A9.947 9.947 0 0 0 8.63 27.966l-.01.005a473.615 473.615 0 0 0 4.989 4.972a2.131 2.131 0 0 1 .638 1.525v6.177a2 2 0 0 0 2 2h15.504a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.281 42.645V21.242A4.281 4.281 0 0 0 24 16.962h0a4.281 4.281 0 0 0-4.281 4.28v21.403"/>`),
		g.Group(children),
	)
}