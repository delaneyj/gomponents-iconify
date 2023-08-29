package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smallbluediamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00B1FF" d="M413.82 266.006c5.412-5.412 5.412-14.267 0-19.678L264.689 97.197c-4.87-4.87-12.84-4.87-17.711 0L96.863 247.311c-4.87 4.87-4.87 12.84 0 17.711l149.131 149.131c5.412 5.412 14.267 5.412 19.678 0L413.82 266.006z"/>`),
		g.Group(children),
	)
}