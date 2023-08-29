package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudVpcEndpoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 27H7a2.006 2.006 0 0 1-2-2V12h2v13h13Z"/><path fill="currentColor" d="m23.4 22l-4-4a3.606 3.606 0 0 0 .6-2a4.012 4.012 0 0 0-4-4a3.606 3.606 0 0 0-2 .6l-4-4V2H2v8h6.6l4 4a3.606 3.606 0 0 0-.6 2a4.012 4.012 0 0 0 4 4a3.606 3.606 0 0 0 2-.6l4 4V30h8v-8ZM8 8H4V4h4Zm8 10a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Zm12 10h-4v-4h4Z"/><path fill="currentColor" d="M25 20h2V7a2.006 2.006 0 0 0-2-2H12v2h13Z"/>`),
		g.Group(children),
	)
}