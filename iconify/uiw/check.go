package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.635 4.243a.795.795 0 0 1 1.12-.023a.788.788 0 0 1 .024 1.117L8.787 16.757a.795.795 0 0 1-1.137.008L.228 9.262a.788.788 0 0 1 .008-1.117a.795.795 0 0 1 1.122.008l6.849 6.924L18.635 4.243Z"/>`),
		g.Group(children),
	)
}