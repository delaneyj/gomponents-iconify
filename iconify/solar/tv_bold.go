package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M22 16v-4c0-2.828 0-4.243-.879-5.121c-.825-.826-2.123-.876-4.621-.879v16c2.498-.003 3.796-.053 4.621-.879c.879-.878.879-2.293.879-5.12Zm-3-5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z" clip-rule="evenodd"/><path d="M15.57 3.488L13.415 6H15v16H8c-2.828 0-4.243 0-5.121-.879C2 20.243 2 18.828 2 16.001v-4c0-2.83 0-4.244.879-5.122C3.757 6 5.172 6 8 6h2.584L8.43 3.488a.75.75 0 0 1 1.138-.976L12 5.348l2.43-2.836a.75.75 0 0 1 1.14.976Z"/></g>`),
		g.Group(children),
	)
}