package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M43 41H23M38.718 5H21L5 41h17.662L38.718 5ZM9.959 29.882h8.028m-4.722-7.412h8.028m-4.519-7.87h8.029"/><path d="M21 5L5 41"/></g>`),
		g.Group(children),
	)
}