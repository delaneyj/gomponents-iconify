package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorSmartphoneLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M11 17H8c-2.828 0-4.243 0-5.121-.879C2 15.243 2 13.828 2 11v-1c0-3.771 0-5.657 1.172-6.828C4.343 2 6.229 2 10 2h5.5c2.335 0 3.502 0 4.386.472a4 4 0 0 1 1.642 1.642C22 4.998 22 6.165 22 8.5"/><path d="M14 15c0-1.886 0-2.828.586-3.414C15.172 11 16.114 11 18 11c1.886 0 2.828 0 3.414.586C22 12.172 22 13.114 22 15v3c0 1.886 0 2.828-.586 3.414C20.828 22 19.886 22 18 22c-1.886 0-2.828 0-3.414-.586C14 20.828 14 19.886 14 18v-3Z"/><path stroke-linecap="round" d="M19 20h-2m-6 2H8m3 0v-5m0-4H2"/></g>`),
		g.Group(children),
	)
}