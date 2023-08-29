package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20.232 10.526c-.585-3.121-.878-4.682-1.989-5.604C17.133 4 15.545 4 12.37 4h-.721c-3.176 0-4.763 0-5.874.922c-1.111.922-1.404 2.483-1.99 5.604c-.822 4.389-1.234 6.583-.034 8.029C4.95 20 7.182 20 11.648 20h.721c4.465 0 6.698 0 7.898-1.445c.696-.84.85-1.93.695-3.555"/><path d="M9.17 8a3.001 3.001 0 0 0 5.66 0"/></g>`),
		g.Group(children),
	)
}