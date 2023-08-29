package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 3a9 9 0 0 1 9 9a1 1 0 0 1-.883.993L20 13h-7v5a1 1 0 0 0 1.993.117L15 18a1 1 0 0 1 2 0a3 3 0 0 1-5.995.176L11 18v-5H4a1 1 0 0 1-.993-.883L3 12a9 9 0 0 1 9-9z"/></g>`),
		g.Group(children),
	)
}