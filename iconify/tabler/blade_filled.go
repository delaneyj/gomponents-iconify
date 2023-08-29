package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BladeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.586 3a2 2 0 0 1 2.828 0l.586.585l.586-.585a2 2 0 0 1 2.7-.117l.128.117L21 5.586a2 2 0 0 1 0 2.828L20.414 9l.586.586a2 2 0 0 1 0 2.828L12.414 21a2 2 0 0 1-2.828 0L9 20.414L8.414 21a2 2 0 0 1-2.828 0L3 18.414a2 2 0 0 1 0-2.828l.585-.587L3 14.414a2 2 0 0 1-.117-2.7L3 11.585zm3.027 4.21a1 1 0 0 0-1.32 1.497l.292.293l-1.068 1.067a2.003 2.003 0 0 0-2.512 1.784L10 12l.005.15c.01.125.03.248.062.367L9 13.585l-.293-.292l-.094-.083a1 1 0 0 0-1.32 1.497l.292.293l-.292.293l-.083.094a1 1 0 0 0 1.497 1.32L9 16.415l.293.292l.094.083a1 1 0 0 0 1.32-1.497L10.415 15l1.069-1.067a2.003 2.003 0 0 0 2.449-2.45L15 10.415l.293.292l.094.083a1 1 0 0 0 1.32-1.497L16.415 9l.292-.293l.083-.094a1 1 0 0 0-1.497-1.32L15 7.585l-.293-.292z"/></g>`),
		g.Group(children),
	)
}