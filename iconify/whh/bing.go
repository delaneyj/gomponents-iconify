package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M511.998 1024q-139 0-257-47t-186.5-128t-68.5-177V32q0-13 9.5-22.5t22.5-9.5h128q13 0 22.5 9.5t9.5 22.5v365q141-77 320-77q139 0 257 47t186.5 128t68.5 177t-68.5 177t-186.5 128t-257 47zm0-576q-87 0-160.5 30t-116.5 81.5t-43 112.5t43 112.5t116.5 81.5t160.5 30t160.5-30t116.5-81.5t43-112.5t-43-112.5t-116.5-81.5t-160.5-30z"/>`),
		g.Group(children),
	)
}