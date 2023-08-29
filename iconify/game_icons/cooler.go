package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cooler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m71.03 41l-11.5 46H452.5L441 41H71.03zM25 105v46h158v-28h146v28h158v-46H25zm176 36v78h110v-78H201zM41 169v276.3L68.82 487H443.2l27.8-41.7V169H329v68H183v-68H41zm196.5 78.1l16.9 30.3h.8l16.9-30.3l15.8 8.8l-24.1 43.2v39.8l33.5-19.8l24.4-42.9l15.6 8.8l-17.3 30.4l.7 1.1l34-.1v18l-48.4.2l-34.3 20.3l34.4 20.4l48.5.3l-.2 18l-33.9-.2l-.7 1.1l17.3 30.4l-15.6 8.8l-24.4-42.9l-33.6-19.9v39.8l24.1 43.2l-15.8 8.8l-16.8-30.2h-1l-16.8 30.2l-15.8-8.8l24.1-43.2v-40.2l-33.5 19.8l-24.4 42.9l-15.6-8.8l17.3-30.4l-.7-1.1l-34 .1v-18l48.4-.2l33.5-19.9l-33.5-19.8l-48.3-.2v-18l34.1.1l.6-1l-17.3-30.5l15.6-8.8l24.5 43l33.3 19.7v-40.3l-24.1-43.2l15.8-8.8z"/>`),
		g.Group(children),
	)
}