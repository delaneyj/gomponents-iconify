package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angularjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.172 24l-9.468-5.244L0 3.984L11.172 0l11.172 3.984l-1.704 14.772zm0-21.348l-6.984 15.66h2.604l1.404-3.504h5.928l1.404 3.504h2.604zm2.04 9.996h-4.08l2.04-4.908z"/>`),
		g.Group(children),
	)
}