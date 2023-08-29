package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDownloadLaptopArrowComputerDownDownloadInternetLaptopNetworkServerUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.91 9.5l-1.3 2.55a1 1 0 0 0 0 1a1 1 0 0 0 .87.47h11a1 1 0 0 0 .87-.47a1 1 0 0 0 0-1L12.09 9.5ZM5 4.5l2 2l2-2m-2 2v-6"/><path d="M2.5 4.63a1 1 0 0 0-.5.87v4h10v-4a1 1 0 0 0-.5-.87"/></g>`),
		g.Group(children),
	)
}