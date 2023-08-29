package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleSwapTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm8 1a3 3 0 1 1-6 0a3 3 0 0 1 6 0ZM2 16.25A2.25 2.25 0 0 1 4.25 14h7.5c.37 0 .72.09 1.027.248l-2.264 2.264a1.75 1.75 0 0 0 0 2.475l1.186 1.187C10.833 20.664 9.641 21 8 21c-6 0-6-4.5-6-4.5v-.25Zm18.28-1.53a.75.75 0 0 0-1.06 1.06L20.44 17h-6.88l1.22-1.22a.75.75 0 0 0-1.06-1.06l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 0 0 1.06-1.06l-1.22-1.22h6.88l-1.22 1.22a.75.75 0 0 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5Z"/>`),
		g.Group(children),
	)
}