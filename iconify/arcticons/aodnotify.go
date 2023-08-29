package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aodnotify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.001 24a14.75 14.75 0 1 1-14.75-14.75h21.633c1.967 0 3.556 1.531 1.967 4.481s-1.361 3.461-5.067 3.461H19.252A6.808 6.808 0 1 0 26.06 24c0-.268.444-.615.817-.615h6.387a.652.652 0 0 1 .737.615Z"/>`),
		g.Group(children),
	)
}