package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneSimCardDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 8.83V20h12V4h-7.17L6 8.83zm5 .19L13 9v4h3l-4 4l-4-4h3V9.02z" opacity=".3"/><path fill="currentColor" d="M18 2h-8L4 8v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 18H6V8.83L10.83 4H18v16z"/><path fill="currentColor" d="m12 17l4-4h-3V9l-2 .02V13H8z"/>`),
		g.Group(children),
	)
}