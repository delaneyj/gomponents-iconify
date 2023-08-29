package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapsTurnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865ZM5 6.5h3s0 0 0 0s4 0 4 4"/><path d="M8.5 9L5 6.5L8.5 4"/></g>`),
		g.Group(children),
	)
}