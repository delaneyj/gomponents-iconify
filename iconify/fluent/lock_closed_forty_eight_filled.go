package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 12a8 8 0 1 1 16 0v2h1.75A6.25 6.25 0 0 1 40 20.25v15.5A6.25 6.25 0 0 1 33.75 42h-19.5A6.25 6.25 0 0 1 8 35.75v-15.5A6.25 6.25 0 0 1 14.25 14H16v-2Zm8-5.5a5.5 5.5 0 0 0-5.5 5.5v2h11v-2A5.5 5.5 0 0 0 24 6.5ZM24 31a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}