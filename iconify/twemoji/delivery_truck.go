package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeliveryTruck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#DD2E44" d="M36 27a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4v-3a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v3z"/><path fill="#FFCC4D" d="m19 13l-.979-1H7.146C4 12 3 14 3 14l-3 5.959V25h19V13z"/><path fill="#55ACEE" d="M9 20H2l2-4s1-2 3-2h2v6z"/><circle cx="9" cy="31" r="4" fill="#292F33"/><circle cx="9" cy="31" r="2" fill="#CCD6DD"/><circle cx="27" cy="31" r="4" fill="#292F33"/><circle cx="27" cy="31" r="2" fill="#CCD6DD"/><path fill="#CCD6DD" d="M32 8H17a4 4 0 0 0-4 4v13h23V12a4 4 0 0 0-4-4z"/>`),
		g.Group(children),
	)
}