package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecureDataOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.8 17v-1.5a2.818 2.818 0 0 0-5.6 0V17a1.29 1.29 0 0 0-1.2 1.2v3.5a1.31 1.31 0 0 0 1.2 1.3h5.5a1.31 1.31 0 0 0 1.3-1.2v-3.5a1.31 1.31 0 0 0-1.2-1.3Zm-4.3-1.5a1.375 1.375 0 0 1 1.5-1.3a1.375 1.375 0 0 1 1.5 1.3V17h-3Z"/><path fill="currentColor" d="M24 11V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h11v-2H2V2h20v9"/><path fill="currentColor" d="M6 4H4v2h2V4zm14 0H8v2h12V4zm-7 12H8v2h5v-2zm-7 0H4v2h2v-2zm4-8H8v2h2V8zm10 0h-8v2h8V8zm-5.192 4H12v2h2.2v-.5a2.26 2.26 0 0 1 .608-1.5ZM10 12H8v2h2v-2z"/>`),
		g.Group(children),
	)
}