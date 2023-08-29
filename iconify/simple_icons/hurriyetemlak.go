package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hurriyetemlak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 16.085L11.994 4.091L0 16.097l3.817 3.812l8.182-8.189l8.189 8.182z"/>`),
		g.Group(children),
	)
}