package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMouseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.435 2.065h-2.87a6.5 6.5 0 0 0-6.5 6.5v6.87a6.5 6.5 0 0 0 6.5 6.5h2.87a6.5 6.5 0 0 0 6.5-6.5v-6.87a6.5 6.5 0 0 0-6.5-6.5Zm-8.37 6.5a5.51 5.51 0 0 1 5.5-5.5h.94v6.44h-6.44Zm13.87 6.87a5.5 5.5 0 0 1-5.5 5.5h-2.87a5.5 5.5 0 0 1-5.5-5.5v-4.93h13.87Zm0-5.93h-6.43v-6.44h.93a5.5 5.5 0 0 1 5.5 5.5Z"/>`),
		g.Group(children),
	)
}