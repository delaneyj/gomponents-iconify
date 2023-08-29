package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DarkcrocTheme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.104 16.104L8.797 8.797m30.406 30.406l-7.488-7.564A10.861 10.861 0 1 1 24 13.136a11.817 11.817 0 0 1 7.804 3.058l7.474-7.32a21.5 21.5 0 1 0-.087 30.34M16.27 31.73l-7.473 7.473"/>`),
		g.Group(children),
	)
}