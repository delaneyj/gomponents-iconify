package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mitid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.38 34.656V13.344h5.796a9.324 9.324 0 0 1 9.324 9.324v2.664a9.324 9.324 0 0 1-9.324 9.324H28.38Zm-23.88 0h0a9.768 9.768 0 0 1 9.768-9.768h0a9.768 9.768 0 0 1 9.768 9.768h0H4.5Z"/><circle cx="14.268" cy="19.116" r="5.772" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}