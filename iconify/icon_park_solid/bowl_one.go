package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBowlOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M24 39c9.389 0 17-7.059 17-17H7c0 9.941 7.611 17 17 17Z"/><path stroke-linejoin="round" d="m18 38l-2 6h16l-2-6"/><path d="M12 10v4m24-4v4M24 5v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBowlOne0)"/>`),
		g.Group(children),
	)
}