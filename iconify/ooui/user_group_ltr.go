package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserGroupLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="6" cy="6" r="3" fill="currentColor"/><circle cx="14" cy="6" r="3" fill="currentColor"/><path fill="currentColor" d="M14 10c3.31 0 6 1.79 6 4v2h-6v-2c0-1.48-1.21-2.77-3-3.46c.88-.35 1.91-.54 3-.54zm-8 0c3.31 0 6 1.79 6 4v2H0v-2c0-2.21 2.69-4 6-4z"/>`),
		g.Group(children),
	)
}