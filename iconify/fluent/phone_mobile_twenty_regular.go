package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneMobileTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M13 2a2 2 0 0 1 1.995 1.85L15 4v12a2 2 0 0 1-1.85 1.995L13 18H7a2 2 0 0 1-1.995-1.85L5 16V4a2 2 0 0 1 1.85-1.995L7 2h6zm0 1H7a1 1 0 0 0-.993.883L6 4v12a1 1 0 0 0 .883.993L7 17h6a1 1 0 0 0 .993-.883L14 16V4a1 1 0 0 0-.883-.993L13 3zm-2 11a.5.5 0 0 1 .09.992L11 15H9a.5.5 0 0 1-.09-.992L9 14h2z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}