package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aerialway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13 5H8V2.6a1 1 0 0 0 .42-.46l5.08-.64a.5.5 0 0 0 0-1l-5.22.65a1 1 0 0 0-.78-.4a1 1 0 0 0-.92.62L1.5 2a.5.5 0 0 0 0 1l5.22-.65c.077.1.172.185.28.25V5H2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1zm-6 6H3V7h4v4zm5 0H8V7h4v4z"/>`),
		g.Group(children),
	)
}