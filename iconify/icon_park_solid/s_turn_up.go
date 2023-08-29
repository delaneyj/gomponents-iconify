package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSTurnUp0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 31v4c0 4-3 7-7 7s-7-3-7-7V16m28 26V13c0-4-3-7-7-7s-7 3-7 7v18"/><path stroke-linecap="round" stroke-linejoin="round" d="m33 37l5 5l5-5"/><circle cx="10" cy="11" r="5" fill="#fff" transform="rotate(-180 10 11)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSTurnUp0)"/>`),
		g.Group(children),
	)
}