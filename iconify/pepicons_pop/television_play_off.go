package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionPlayOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 8v7a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1Zm-2 7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.875 5.78a1 1 0 0 1-.156-1.405l2-2.5a1 1 0 0 1 1.562 1.25l-2 2.5a1 1 0 0 1-1.406.156Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.125 5.78a1 1 0 0 0 .156-1.405l-2-2.5a1 1 0 1 0-1.562 1.25l2 2.5a1 1 0 0 0 1.406.156ZM12 11.5l-3.25 2v-4l3.25 2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.5 11.5a.5.5 0 0 1-.238.426l-3.25 2a.5.5 0 0 1-.762-.426v-4a.5.5 0 0 1 .762-.426l3.25 2a.5.5 0 0 1 .238.426Zm-3.25-1.105v2.21l1.796-1.105l-1.796-1.105Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}