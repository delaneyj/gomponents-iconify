package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTransferUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h8V0H0zm4 2L1 5h2v3h2V5h2L4 2z"/>`),
		g.Group(children),
	)
}