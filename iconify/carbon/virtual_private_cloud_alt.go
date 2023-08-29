package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualPrivateCloudAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.414 22L10 8.586V2H2v8h6.586L22 23.414V30h8v-8ZM8 8H4V4h4Zm20 20h-4v-4h4Z"/><path fill="currentColor" d="M30 6a3.991 3.991 0 0 0-7.858-1H13v2h9.142A3.994 3.994 0 0 0 25 9.858V19h2V9.858A3.996 3.996 0 0 0 30 6zm-4 2a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zm-7 17H9.858A3.994 3.994 0 0 0 7 22.142V13H5v9.142A3.991 3.991 0 1 0 9.858 27H19zM6 28a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}