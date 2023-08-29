package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gtk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.01 23.773V14.45l-7.584 2.245Zm0-13.87L.91 3.828l.502 12.526l7.597-2.249ZM9.57 24l12.353-5.146l-8.285-5.775l-4.068 1.204ZM23.09 5.815l-9.257 2.849v4.148l8.237 5.741ZM9.57 9.975v3.964l3.932-1.164v-4.01Zm-.228-.52l4.16-1.28V0L1.231 3.37ZM22.715 5.34L13.833.052v8.021Z"/>`),
		g.Group(children),
	)
}