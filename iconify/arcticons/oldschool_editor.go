package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OldschoolEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.645 7.343H42.5l-12.145 24.66H5.5l12.145-24.66Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.467 7.353v33.304H18.063V31.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.459 37.101l-18.442.004L39.46 37.1Zm-18.442-2.706l18.443.004l-18.443-.004Zm9.573-2.705l8.87.003l-8.87-.003Zm1.352-2.706l7.518.004l-7.518-.004Zm1.353-2.706l6.165.004l-6.165-.004Zm1.353-2.705l4.812.003l-4.812-.003ZM36 20.867l3.46.003l-3.46-.003Zm1.353-2.706l2.107.004l-2.107-.004Zm2.107-2.702h-.85"/>`),
		g.Group(children),
	)
}