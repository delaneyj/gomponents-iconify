package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IpOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.26 12A8.243 8.243 0 0 1 4 10a8.243 8.243 0 0 1 .26-2h3.38a16.513 16.513 0 0 0-.14 2a16.514 16.514 0 0 0 .14 2h2.02a14.71 14.71 0 0 1-.16-2a14.581 14.581 0 0 1 .16-2h4.68a14.59 14.59 0 0 1 .16 2a14.72 14.72 0 0 1-.16 2h2.02a16.512 16.512 0 0 0 .14-2a16.511 16.511 0 0 0-.14-2h3.38a8.24 8.24 0 0 1 .26 2a8.24 8.24 0 0 1-.26 2L22 13v-3a10 10 0 1 0-20 0v3Zm14.66-6h-2.95a15.65 15.65 0 0 0-1.38-3.56A8.03 8.03 0 0 1 18.92 6ZM12 2.04A14.086 14.086 0 0 1 13.91 6h-3.82A14.086 14.086 0 0 1 12 2.04Zm-2.59.4A15.648 15.648 0 0 0 8.03 6H5.08a7.987 7.987 0 0 1 4.33-3.56Z"/><path fill="currentColor" d="M21 14v8H3v-8h18m0-2H3a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2Z"/><path fill="currentColor" d="M7 15h2v6H7zm9 0h-4v6h2v-2h2a1.473 1.473 0 0 0 1.5-1.5v-1A1.473 1.473 0 0 0 16 15Zm0 3h-2v-2h2Z"/>`),
		g.Group(children),
	)
}