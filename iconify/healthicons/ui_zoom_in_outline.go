package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiZoomInOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M28 16v4h4v2h-4v4h-2v-4h-4v-2h4v-4h2Z"/><path fill-rule="evenodd" d="M42 21c0 8.284-6.716 15-15 15c-3.782 0-7.238-1.4-9.876-3.71l-1.828 1.828l.004 2.825l-6.472 6.471l-4.242-4.242l6.524-6.524l2.707.12l1.893-1.892A14.943 14.943 0 0 1 12 21c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15Zm-2 0c0 7.18-5.82 13-13 13s-13-5.82-13-13S19.82 8 27 8s13 5.82 13 13ZM7.414 39.172l1.414 1.414l4.47-4.47l-.001-1.368l-1.397-.063l-4.486 4.487Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}