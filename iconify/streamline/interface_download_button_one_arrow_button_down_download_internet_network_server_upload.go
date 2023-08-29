package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDownloadButtonOneArrowButtonDownDownloadInternetNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 4.75h1a.5.5 0 0 1 .5.5v7.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5v-7.5a.5.5 0 0 1 .5-.5h1m3.5-4v8"/><path d="m5 6.75l2 2l2-2"/></g>`),
		g.Group(children),
	)
}