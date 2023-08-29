package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundPushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d0d0" d="M42.6 29.4L2 62l32.6-40.6z"/><circle cx="45" cy="19" r="17" fill="#ed4c5c"/>`),
		g.Group(children),
	)
}