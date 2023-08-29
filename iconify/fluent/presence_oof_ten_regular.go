package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceOofTenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.349 3.85a.5.5 0 1 0-.708-.706l-1.497 1.5a.5.5 0 0 0 0 .707l1.497 1.5a.5.5 0 1 0 .708-.707l-.646-.646h1.8a.5.5 0 1 0 0-1h-1.8l.646-.647ZM4.998 0a4.998 4.998 0 1 0 0 9.995a4.998 4.998 0 0 0 0-9.995ZM1 4.998a3.998 3.998 0 1 1 7.995 0a3.998 3.998 0 0 1-7.995 0Z"/>`),
		g.Group(children),
	)
}