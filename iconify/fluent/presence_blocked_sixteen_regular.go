package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceBlockedSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0Zm-2 0c0-1.296-.41-2.496-1.11-3.477l-8.367 8.368A6 6 0 0 0 14 8Zm-2.524-4.89a6 6 0 0 0-8.367 8.367l8.368-8.368Z"/>`),
		g.Group(children),
	)
}