package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineglassBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 4.895C5 3.848 5.848 3 6.895 3h10.21C18.152 3 19 3.848 19 4.895V8A7 7 0 1 1 5 8V4.895Z" opacity=".5"/><path fill-rule="evenodd" d="M11.25 20.25h-3a.75.75 0 0 0 0 1.5h7.5a.75.75 0 0 0 0-1.5h-4.5ZM12 15a7.002 7.002 0 0 0 6.925-5.97c-.402.157-2.332.886-3.694.963c-1.327.075-2.28-.459-3.23-.993c-.952-.534-1.904-1.068-3.232-.993c-1.371.078-3.318.816-3.702.966A7.001 7.001 0 0 0 12 15Z" clip-rule="evenodd"/><path d="M12.75 14.96a7.087 7.087 0 0 1-1.5 0v5.29h1.5v-5.29Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}