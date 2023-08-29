package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.44 7.937H17.3l-1.21-4.51a.508.508 0 0 0-.61-.35a.489.489 0 0 0-.35.61l1.14 4.25H7.74l1.14-4.25a.5.5 0 0 0-.36-.61a.513.513 0 0 0-.61.35l-1.2 4.51H4.56a1.5 1.5 0 0 0-.32 2.96l.74 7.77a2.492 2.492 0 0 0 2.49 2.27h9.06a2.492 2.492 0 0 0 2.49-2.27l.74-7.77a1.5 1.5 0 0 0-.32-2.96Zm-1.41 10.64a1.5 1.5 0 0 1-1.5 1.36H7.47a1.5 1.5 0 0 1-1.5-1.36l-.72-7.64h13.5Zm1.41-8.64H4.56a.508.508 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h14.88a.5.5 0 0 1 .5.5a.508.508 0 0 1-.5.5Z"/><path fill="currentColor" d="M9.5 17.432a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v3a.5.5 0 0 1-.5.5Zm5 0a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v3a.5.5 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}