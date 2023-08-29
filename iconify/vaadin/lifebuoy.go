package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifebuoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM4 8c0-2.2 1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4s-4-1.8-4-4zm8.6 1.8c.3-.5.4-1.2.4-1.8s-.1-1.3-.4-1.8l1.5-1.5c.6 1 .9 2.1.9 3.3s-.3 2.3-.8 3.3l-1.6-1.5zm-1.3-8L9.8 3.4C9.3 3.1 8.6 3 8 3s-1.3.1-1.8.4L4.7 1.8C5.7 1.3 6.8 1 8 1s2.3.3 3.3.8zM1.8 4.7l1.5 1.5C3.1 6.7 3 7.4 3 8s.1 1.3.4 1.8l-1.5 1.5C1.3 10.3 1 9.2 1 8s.3-2.3.8-3.3zm2.9 9.5l1.5-1.5c.5.2 1.2.3 1.8.3s1.3-.1 1.8-.4l1.5 1.5c-1 .6-2.1.9-3.3.9s-2.3-.3-3.3-.8z"/>`),
		g.Group(children),
	)
}