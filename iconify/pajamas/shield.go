package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 5.4V3.132l4.75-1.357v11.608l-1.782-1.528A8.5 8.5 0 0 1 2.5 5.401Zm6.25 7.982l1.782-1.528A8.5 8.5 0 0 0 13.5 5.401V3.13L8.75 1.774v11.608ZM1 2l7-2l7 2v3.4a10 10 0 0 1-3.492 7.593L8 16l-3.508-3.007A10 10 0 0 1 1 5.401V2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}