package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 45v604c0 17 14 31 31 31h604c17 0 31-14 31-31V45c0-17-14-31-31-31H31C14 14 0 28 0 45z"/>`),
		g.Group(children),
	)
}