package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAvailableTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0ZM3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Zm14.06-3.56a1.5 1.5 0 0 1 0 2.12l-5.5 5.5a1.5 1.5 0 0 1-2.12 0l-2-2a1.5 1.5 0 0 1 2.12-2.12l.94.939l4.44-4.44a1.5 1.5 0 0 1 2.12 0Z"/>`),
		g.Group(children),
	)
}