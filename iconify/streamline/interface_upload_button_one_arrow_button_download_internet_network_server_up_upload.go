package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUploadButtonOneArrowButtonDownloadInternetNetworkServerUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 5h1a.5.5 0 0 1 .5.5V13a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 2 13V5.5a.5.5 0 0 1 .5-.5h1M7 7.5v-7m-2 2l2-2l2 2"/>`),
		g.Group(children),
	)
}