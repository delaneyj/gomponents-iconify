package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkSearchTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.596 7.303a3.5 3.5 0 1 1 .707-.707l2.55 2.55a.5.5 0 0 1-.707.708l-2.55-2.55ZM16 4.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Zm0 4.621V17.5a.5.5 0 0 1-.794.404L10 14.118l-5.206 3.786A.5.5 0 0 1 4 17.5v-13A2.5 2.5 0 0 1 6.5 2h3.258a4.484 4.484 0 0 0-.502 1H6.5A1.5 1.5 0 0 0 5 4.5v12.018l4.706-3.422a.5.5 0 0 1 .588 0L15 16.518V8.744c.15-.053.297-.114.44-.183l.56.56Z"/>`),
		g.Group(children),
	)
}