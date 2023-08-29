package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareScreenStartTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 4A2.25 2.25 0 0 0 2 6.25v11.5A2.25 2.25 0 0 0 4.25 20h15.5A2.25 2.25 0 0 0 22 17.75V6.25A2.25 2.25 0 0 0 19.75 4H4.25Zm8.28 3.465l3.255 3.255a.75.75 0 1 1-1.06 1.06L12.75 9.806v6.447a.75.75 0 1 1-1.5 0V9.808L9.28 11.78a.75.75 0 1 1-1.06-1.06l3.25-3.254a.75.75 0 0 1 1.06 0Z"/>`),
		g.Group(children),
	)
}