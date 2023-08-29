package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuteNotification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 20a2.5 2.5 0 0 1-5 0M.5.5l23 23m-6-6h5v-.25l-.848-.908a12 12 0 0 1-3.134-6.7l-.336-2.684A6.23 6.23 0 0 0 6.011 6.01M14 17.5H1.5v-.25l.848-.908a12 12 0 0 0 3.134-6.7l.074-.586"/>`),
		g.Group(children),
	)
}