package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M28 16v4h4v2h-4v4h-2v-4h-4v-2h4v-4h2Z"/><path fill-rule="evenodd" d="M42 21c0 8.284-6.716 15-15 15a14.942 14.942 0 0 1-9.784-3.63l-1.346 1.346l-.158 1.4L8.828 42L6 39.17l6.927-6.926l1.705-.12l1.158-1.158A14.944 14.944 0 0 1 12 21c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15Zm-2 0c0 7.18-5.82 13-13 13s-13-5.82-13-13S19.82 8 27 8s13 5.82 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}