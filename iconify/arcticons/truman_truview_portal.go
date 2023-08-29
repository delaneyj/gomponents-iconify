package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrumanTruviewPortal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.23 17.756l9.542 5.881V43.5H19.23V17.756ZM41.064 6.41v8.619H10.466L41.064 6.41ZM6.936 4.5h21.927L6.936 14.457V4.5Z"/>`),
		g.Group(children),
	)
}