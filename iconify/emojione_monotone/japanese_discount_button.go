package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseDiscountButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M19.003 23.592h3.888v-2.098h4.154v2.098h3.889v-2.875H19.003zm2.21 17.941h7.585v3.301h-7.585z"/><path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10M33.105 49.611h-4.307v-1.32h-7.585V50h-4.077V38.078h15.969v11.533m2.022-25.709h-1.982v2.951h-6.099v1.631h5.564v3.067h-5.564v1.671h7.662v3.535H15.116v-3.535h7.775v-1.671H17.25v-3.067h5.641v-1.631h-6.023v-2.951H15v-6.68h7.776V14h4.346v3.223h8.004v6.679zm1.6-6.021h4.309v21.711h-4.309V17.881M49 44.406c0 2.564-.457 3.848-1.982 4.584c-1.486.777-3.85.932-6.975.932c-.189-1.32-.84-3.455-1.486-4.777c2.211.117 4.385.078 5.146.078c.646 0 .914-.232.914-.854V14.465H49v29.941"/>`),
		g.Group(children),
	)
}