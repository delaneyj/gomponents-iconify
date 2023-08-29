package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDGlasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M19.6 34.5h8.8"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.303 27.39L9.958 9.563L14.949 7m30.855 20.633L38.05 9.562L33.059 7"/><path d="M3.4 26.8h15.317a1.294 1.294 0 0 1 1.29 1.396l-.927 11.712a1.294 1.294 0 0 1-1.29 1.192H4.328a1.294 1.294 0 0 1-1.29-1.192l-.927-11.712A1.294 1.294 0 0 1 3.4 26.8Zm25.884 0H44.6a1.294 1.294 0 0 1 1.29 1.396l-.928 11.712a1.294 1.294 0 0 1-1.29 1.192h-13.46a1.294 1.294 0 0 1-1.29-1.192l-.928-11.712a1.294 1.294 0 0 1 1.29-1.396Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}