package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.5 5.777v6.32l3-1.874V3.902l-3 1.875Zm4-1.875v6.32l3 1.876V5.777l-3-1.875ZM6 11.09l-3.735 2.334L1.5 13V5.5l.235-.424l4-2.5h.53L10 4.91l3.735-2.334L14.5 3v7.5l-.235.424l-4 2.5h-.53L6 11.09Zm4.5-5.313v6.32l3-1.874V3.902l-3 1.875Z"/>`),
		g.Group(children),
	)
}