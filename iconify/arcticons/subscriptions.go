package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subscriptions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 14.039a.95.95 0 0 1 .949-.95H41.55a.95.95 0 0 1 .949.95v19.923a.949.949 0 0 1-.949.948H6.45a.949.949 0 0 1-.949-.948V14.038Zm3.32 12.807H24M8.82 29.692h7.59"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.18 26.846a1.897 1.897 0 1 1-3.795 0a1.897 1.897 0 0 1 3.794 0Zm-3.592-.949H31.59a.949.949 0 1 0 0 1.898h3.998M8.346 16.885a.95.95 0 0 1 .949-.95h19.923a.95.95 0 0 1 .949.95v2.846a.949.949 0 0 1-.949.948H9.295a.949.949 0 0 1-.949-.948v-2.846ZM7.872 34.91v1.949c0 .786.637 1.423 1.423 1.423h29.41c.786 0 1.423-.637 1.423-1.423V34.91M7.872 13.09v-1.95c0-.786.637-1.423 1.423-1.423h29.41c.786 0 1.423.637 1.423 1.423v1.949"/>`),
		g.Group(children),
	)
}