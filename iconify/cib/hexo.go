package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 .011L2.093 8.074v15.995L16 31.996l13.907-8.068V7.939zm6.396 22.806l-1.251.693l-1.249-.693v-5.563h-7.792v5.563l-1.249.693l-1.251-.693V9.188l1.251-.699l1.249.699v5.561h7.792V9.188l1.249-.699l1.251.699z"/>`),
		g.Group(children),
	)
}