package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteForeverTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M16.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zM12 1.5A3.5 3.5 0 0 1 15.5 5h5a1 1 0 0 1 0 2h-.845l-.451 4.587A6.5 6.5 0 0 0 11.81 22H8.312a2.75 2.75 0 0 1-2.737-2.48L4.345 7H3.5a1 1 0 1 1 0-2h5A3.5 3.5 0 0 1 12 1.5zm1.716 13.089l-.07.057l-.057.07a.5.5 0 0 0 0 .568l.057.07l2.147 2.146l-2.147 2.146l-.057.07a.5.5 0 0 0 0 .568l.057.07l.07.057a.5.5 0 0 0 .568 0l.07-.057l2.146-2.147l2.146 2.147l.07.057a.5.5 0 0 0 .568 0l.07-.057l.057-.07a.5.5 0 0 0 0-.568l-.057-.07l-2.147-2.146l2.147-2.146l.057-.07a.5.5 0 0 0 0-.568l-.057-.07l-.07-.057a.5.5 0 0 0-.568 0l-.07.057l-2.146 2.147l-2.146-2.147l-.07-.057a.5.5 0 0 0-.492-.044l-.076.044zM12 3.5A1.5 1.5 0 0 0 10.5 5h3A1.5 1.5 0 0 0 12 3.5z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}