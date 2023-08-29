package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceBlockedTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0Zm-1.5 0a4.48 4.48 0 0 0-.832-2.607L3.393 9.668A4.5 4.5 0 0 0 10.5 6ZM8.607 2.332a4.5 4.5 0 0 0-6.275 6.275l6.275-6.275Z"/>`),
		g.Group(children),
	)
}