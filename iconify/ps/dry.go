package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M429 43q0-18-12.5-30.5T387 0H45Q28 0 15.5 12.5T3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43zm-42 426H45V43h342v426z"/>`),
		g.Group(children),
	)
}