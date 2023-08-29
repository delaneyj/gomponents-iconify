package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.9 4.4c-.9 5-5.9 9.3-7.4 10.3s-2.9-.4-3.4-1.4C4.6 12 2.9 5.7 2.4 5.1C2 4.6.6 5.7.6 5.7L0 4.8s2.7-3.3 4.8-3.7C7 .7 7 4.5 7.5 6.6c.5 2 .9 3.2 1.3 3.2s1.3-1.1 2.2-2.9c.9-1.7 0-3.3-1.9-2.2c.8-4.3 7.7-5.4 6.8-.3z"/>`),
		g.Group(children),
	)
}