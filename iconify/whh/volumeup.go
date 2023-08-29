package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volumeup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m874 873l-45-45q62-62 96.5-143.5t34.5-173t-34.5-173T829 195l45-46q71 71 110.5 164.5t39.5 198t-39.5 198T874 873zM761 760l-45-45q84-84 84-203.5T716 308l45-45q49 48 76 112.5t27 136t-27 136T761 760zm-196 251L320 768V256L565 12q30-30 75 16v968q-45 45-75 15zM256 768H64q-27 0-45.5-19T0 704V319q0-26 18.5-45T64 255h192v513z"/>`),
		g.Group(children),
	)
}