package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projectsend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M122 281L592 10q27-15 56.5-7T693 37l321 555q13 23 9 48q-307 0-553-100T122 281zm847 422l-538 311q-26 15-55.5 7.5T330 987L0 448q142 115 402.5 184T969 703z"/>`),
		g.Group(children),
	)
}