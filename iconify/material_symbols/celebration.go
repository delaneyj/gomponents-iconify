package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Celebration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L7 8l9 9l-14 5Zm12.55-9.45L13.5 11.5l5.6-5.6q.8-.8 1.925-.8t1.925.8l.6.6l-1.05 1.05l-.6-.6q-.35-.35-.875-.35t-.875.35l-5.6 5.6Zm-4-4L9.5 7.5l.6-.6q.35-.35.35-.85t-.35-.85l-.65-.65L10.5 3.5l.65.65q.8.8.8 1.9t-.8 1.9l-.6.6Zm2 2L11.5 9.5l3.6-3.6q.35-.35.35-.875t-.35-.875l-1.6-1.6l1.05-1.05l1.6 1.6q.8.8.8 1.925t-.8 1.925l-3.6 3.6Zm4 4L15.5 13.5l1.6-1.6q.8-.8 1.925-.8t1.925.8l1.6 1.6l-1.05 1.05l-1.6-1.6q-.35-.35-.875-.35t-.875.35l-1.6 1.6Z"/>`),
		g.Group(children),
	)
}