package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOpenOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 18V9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v9m-21 8l12-13m-7 8l4 4m1-9l4 4M6 30v9a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-9"/><circle cx="18" cy="30" r="5" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOpenOne0)"/>`),
		g.Group(children),
	)
}