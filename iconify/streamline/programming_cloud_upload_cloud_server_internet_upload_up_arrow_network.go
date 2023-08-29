package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingCloudUploadCloudServerInternetUploadUpArrowNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 7.5v6m-2-4l2-2l2 2"/><path d="M12 8.1a3 3 0 1 0-3.65-4.69A4 4 0 0 0 .5 4.5A4 4 0 0 0 1.38 7"/></g>`),
		g.Group(children),
	)
}