package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stackoverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M245 395V250h34v179H0V251l32 1l-1 143h214zM52 335h167v35H52v-35zm5-63l168 16l-4 36l-168-16zm15-73l163 46l-10 35l-163-46zm40-82l144 87l-19 32l-144-87zm150 81L164 61l30-21l98 137zM272 9l36-6l28 166l-36 6z"/>`),
		g.Group(children),
	)
}