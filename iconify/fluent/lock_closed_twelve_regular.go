package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4V3a2 2 0 1 1 4 0v1h.5A1.5 1.5 0 0 1 10 5.5v4A1.5 1.5 0 0 1 8.5 11h-5A1.5 1.5 0 0 1 2 9.5v-4A1.5 1.5 0 0 1 3.5 4H4Zm1-1v1h2V3a1 1 0 0 0-2 0ZM3.5 5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-5Z"/>`),
		g.Group(children),
	)
}