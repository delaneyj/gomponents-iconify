package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GestureTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 18a1 1 0 1 1 0 2a1 1 0 0 1 0-2ZM7 4h10a1 1 0 0 1 .117 1.993L17 6h-4.649l8.01 3.102c.77.298.855 1.33.195 1.764l-.11.064l-14 6.965a1 1 0 0 1-.993-1.732l.102-.058l11.97-5.956L6.64 5.933c-.994-.386-.766-1.821.241-1.927L7 4h10H7Zm13 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/>`),
		g.Group(children),
	)
}