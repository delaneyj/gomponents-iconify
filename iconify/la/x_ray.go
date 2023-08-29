package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XRay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm8 1v1h-3v2h3v1h-5v2h5v1h-4v2h4v2.563c-.523-.27-1.113-.563-1.5-.563a1.5 1.5 0 0 0 0 3c.379 0 1.672 1 2.5 1c.828 0 2.121-1 2.5-1a1.5 1.5 0 0 0 0-3c-.387 0-.977.293-1.5.563V8zm3 1v2h2V9zm0 3v2h4v-2zm0 3v2h3v-2z"/>`),
		g.Group(children),
	)
}