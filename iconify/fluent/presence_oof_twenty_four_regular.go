package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceOofTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.562 9.057a1.5 1.5 0 0 0-2.121-2.122L6.437 10.94a1.5 1.5 0 0 0 0 2.12l4.004 4.005a1.5 1.5 0 0 0 2.121-2.121L11.12 13.5h5.38a1.5 1.5 0 1 0 0-3h-5.38l1.443-1.443ZM12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0ZM3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Z"/>`),
		g.Group(children),
	)
}