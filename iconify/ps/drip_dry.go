package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DripDry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43q0-18-12.5-30.5T387 0H45Q28 0 15.5 12.5T3 43zm384 426H45V43h342v426zm-256-64q21 0 21-21V128q0-21-21-21q-22 0-22 21v256q0 21 22 21zm85 0q21 0 21-21V128q0-21-21-21t-21 21v256q0 21 21 21zm85 0q22 0 22-21V128q0-21-22-21q-21 0-21 21v256q0 21 21 21z"/>`),
		g.Group(children),
	)
}