package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.6l-.3-.3l-6-6l-.3-.3H6zm2 2h10v6h6v16H8V5zm12 1.4L22.6 9H20V6.4zM16 13l-4 4h3v5h2v-5h3l-4-4zm-4 10v2h8v-2h-8z"/>`),
		g.Group(children),
	)
}