package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.097 8.038v-5.19c.038-.343.227-.563.568-.662c.341-.1.678-.013 1.012.26l8.086 6.838a.924.924 0 0 1 .334.716c0 .28-.111.52-.334.718l-8.14 6.884c-.333.23-.66.299-.98.205c-.32-.093-.502-.298-.546-.613v-5.238l-6.614 5.598c-.33.264-.68.349-1.046.253c-.366-.095-.544-.3-.534-.613V2.93c-.01-.329.14-.566.451-.712c.311-.146.668-.089 1.07.172l6.673 5.647Z"/>`),
		g.Group(children),
	)
}