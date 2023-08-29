package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pharmeasy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.722 32.211l6.9-6.9L33.911 35.6l-6.9 6.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.098 22.587l13.89-13.89c4.263-4.262 11.173-4.262 15.435 0h0c4.262 4.262 4.262 11.172 0 15.434l4.479 4.478"/>`),
		g.Group(children),
	)
}