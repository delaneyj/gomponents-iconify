package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.6l-.3-.3l-6-6l-.3-.3H6zm2 2h10v6h6v16H8V5zm12 1.4L22.6 9H20V6.4zM15 13v5h-3l4 4l4-4h-3v-5h-2zm-3 10v2h8v-2h-8z"/>`),
		g.Group(children),
	)
}