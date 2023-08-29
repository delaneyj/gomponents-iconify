package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotStep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.5 15.675l-1.4-1.4L21.575 9.8L23 11.2l-4.5 4.475Zm-2.1-2.125L9.3 6.475L13.8 2l7.075 7.075L16.4 13.55Zm3.375 9.05l-5.95-5.925L10.5 20H1v-2.625q0-.65.363-1.15T2.3 15.5q.425-.175.938-.425t1.037-.6L5.65 15.85Q5.8 16 6 16t.35-.15q.15-.15.15-.35t-.15-.35l-1.275-1.3l.387-.388q.188-.187.388-.412l1.275 1.275q.15.15.35.15t.35-.15q.15-.15.15-.35t-.15-.35l-1.4-1.4q.15-.275.263-.55t.187-.575l1.7 1.725q.15.15.35.15t.35-.15q.15-.15.162-.35t-.137-.35l-7.9-7.9L2.8 2.8l18.375 18.4l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}