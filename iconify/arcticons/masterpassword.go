package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Masterpassword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.94 19.56A4.44 4.44 0 1 1 4.5 24a4.44 4.44 0 0 1 4.44-4.44Zm12.37 0A4.44 4.44 0 1 1 16.87 24a4.44 4.44 0 0 1 4.44-4.44Zm12.37 0A4.44 4.44 0 1 1 29.24 24a4.44 4.44 0 0 1 4.44-4.44Zm9.82-7.09v23.06"/>`),
		g.Group(children),
	)
}