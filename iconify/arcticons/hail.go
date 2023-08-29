package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 23.835l10.06 16.92h18.696L43.5 23.847"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.332 40.595l19.152-33.35H14.388l18.834 33.422"/>`),
		g.Group(children),
	)
}