package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.548 4.078h16.9a.5.5 0 0 0 0-1h-16.9a.5.5 0 0 0 0 1ZM9 20.922H6.565a2.5 2.5 0 0 1-2.5-2.5V7.582a2.5 2.5 0 0 1 2.5-2.5H9a2.5 2.5 0 0 1 2.5 2.5v10.84a2.5 2.5 0 0 1-2.5 2.5ZM6.565 6.082a1.5 1.5 0 0 0-1.5 1.5v10.84a1.5 1.5 0 0 0 1.5 1.5H9a1.5 1.5 0 0 0 1.5-1.5V7.582a1.5 1.5 0 0 0-1.5-1.5Zm10.873 9.869H15a2.5 2.5 0 0 1-2.5-2.5V7.582a2.5 2.5 0 0 1 2.5-2.5h2.435a2.5 2.5 0 0 1 2.5 2.5v5.869a2.5 2.5 0 0 1-2.497 2.5ZM15 6.082a1.5 1.5 0 0 0-1.5 1.5v5.869a1.5 1.5 0 0 0 1.5 1.5h2.435a1.5 1.5 0 0 0 1.5-1.5V7.582a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}