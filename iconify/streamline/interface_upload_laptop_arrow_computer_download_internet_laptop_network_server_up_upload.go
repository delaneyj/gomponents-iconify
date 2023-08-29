package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUploadLaptopArrowComputerDownloadInternetLaptopNetworkServerUpUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.91 9.5l-1.3 2.55a1 1 0 0 0 0 1a1 1 0 0 0 .87.47h11a1 1 0 0 0 .87-.47a1 1 0 0 0 0-1L12.09 9.5ZM5 2.5l2-2l2 2m-2-2v6"/><path d="M3 4.5a1 1 0 0 0-1 1v4h10v-4a1 1 0 0 0-1-1"/></g>`),
		g.Group(children),
	)
}