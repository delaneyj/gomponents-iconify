package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iconify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M24 3.5c-2.501 0-13.5 10.942-15.878 18.888h.035a16.753 16.753 0 0 0-.915 5.357C7.242 36.998 14.745 44.5 24 44.5c9.255 0 16.758-7.502 16.758-16.755a16.74 16.74 0 0 0-.882-5.357h.004C37.502 14.442 26.5 3.5 24 3.5Z"/><circle cx="24" cy="18.023" r="3.852" stroke-linecap="round"/><path d="M18.403 23.785c2.297.096 3.044 3.357 5.598 3.788c2.553.431 3.834-.772 3.834-.772A5.7 5.7 0 0 1 24 36.719c-3.148 0-5.7-2.552-5.7-5.676l.102-7.258Z"/></g>`),
		g.Group(children),
	)
}