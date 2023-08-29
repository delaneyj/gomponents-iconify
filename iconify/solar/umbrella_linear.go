package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M8.75 19a.75.75 0 0 0-1.5 0h1.5Zm2.5-7v8h1.5v-8h-1.5Zm-2.5 8v-1h-1.5v1h1.5ZM12 12.75h9.524v-1.5H12v1.5Zm.476-11.5h-.952v1.5h.952v-1.5Zm-10 11.5H12v-1.5H2.476v1.5ZM1.25 11.524c0 .677.549 1.226 1.226 1.226v-1.5c.151 0 .274.123.274.274h-1.5Zm21.5 0c0-5.674-4.6-10.274-10.274-10.274v1.5a8.774 8.774 0 0 1 8.774 8.774h1.5Zm-1.226 1.226c.677 0 1.226-.549 1.226-1.226h-1.5c0-.151.123-.274.274-.274v1.5ZM10 21.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 10 22.75v-1.5Zm1.524-20C5.85 1.25 1.25 5.85 1.25 11.524h1.5a8.774 8.774 0 0 1 8.774-8.774v-1.5ZM11.25 20c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 12.75 20h-1.5Z"/><path stroke="currentColor" stroke-width="1.5" d="M16.915 12C16.634 9.289 15.555 2 12 2C8.446 2 7.365 9.289 7.086 12"/></g>`),
		g.Group(children),
	)
}