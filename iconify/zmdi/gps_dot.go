package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M234.5 155q35.5 0 60.5 25t25 60t-25 60t-60.5 25t-60.5-25t-25-60t25-60t60.5-25zM425 219h44v42h-44q-7 67-54.5 114.5T256 431v44h-43v-44q-66-8-114-55.5T44 261H0v-42h44q7-67 55-114.5T213 49V5h43v44q67 8 114.5 55.5T425 219zM235 389q62 0 105.5-43.5T384 240t-43.5-105.5T235 91t-106 43.5T85 240t44 105.5T235 389z"/>`),
		g.Group(children),
	)
}