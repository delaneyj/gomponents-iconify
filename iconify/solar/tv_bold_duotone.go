package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16 6H8c-2.828 0-4.243 0-5.121.879C2 7.757 2 9.172 2 12v4c0 2.828 0 4.243.879 5.121C3.757 22 5.172 22 8 22h8V6Z" clip-rule="evenodd"/><path d="M22 12v4c0 2.828 0 4.242-.879 5.121c-.825.825-2.123.876-4.621.879H16V6h.5c2.498.003 3.796.053 4.621.878C22 7.758 22 9.171 22 12Zm-8.584-6l2.153-2.512a.75.75 0 0 0-1.138-.976L12 5.348L9.57 2.512a.75.75 0 0 0-1.14.976L10.585 6h2.832Z" opacity=".5"/><path d="M19 11a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/></g>`),
		g.Group(children),
	)
}