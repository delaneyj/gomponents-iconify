package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DryFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M45 0Q28 0 15.5 12.5T3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43q0-18-12.5-30.5T387 0H45zm342 469H45V43h342v426zM88 277h256q21 0 21-21t-21-21H88q-21 0-21 21t21 21z"/>`),
		g.Group(children),
	)
}