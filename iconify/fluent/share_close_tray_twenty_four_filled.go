package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareCloseTrayTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 20A2.25 2.25 0 0 1 2 17.75V6.25A2.25 2.25 0 0 1 4.25 4h15.5A2.25 2.25 0 0 1 22 6.25v11.5A2.25 2.25 0 0 1 19.75 20H4.25Zm8.28-3.465l3.255-3.255a.75.75 0 1 0-1.06-1.06l-1.975 1.974V7.747a.75.75 0 0 0-1.5 0v6.445L9.28 12.22a.75.75 0 1 0-1.06 1.06l3.25 3.255a.75.75 0 0 0 1.06 0Z"/>`),
		g.Group(children),
	)
}