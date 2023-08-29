package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromecastActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2 20.01l.01-.011M15 20h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v1m0 9c2 .5 3.5 2 4 4m-4-8c4 .5 7.5 4 8 8"/><path fill="currentColor" fill-rule="evenodd" d="M5.002 7.63a.6.6 0 0 1 .6-.6h12.804a.6.6 0 0 1 .6.6v8.832a.6.6 0 0 1-.6.6H13.44a.617.617 0 0 1-.556-.355c-.422-.892-1.622-3.26-3.07-4.707c-1.42-1.419-3.572-2.444-4.435-2.82a.624.624 0 0 1-.378-.569v-.98Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}