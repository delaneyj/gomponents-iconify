package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.932 9.009V16H15v4.009h2V16h3.932v-2H17V7.009H9.932V3h-2v4.009H4v2h3.932Zm2 0V14H15V9.009H9.932Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}