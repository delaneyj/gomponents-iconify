package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 28v-2h2v2zm-6-4v-2h2v2zm0 6h4v-2h-2v-2h-2v4zm8-4v-4h2v4zm2 0h2v4h-4v-2h2v-2zm-2-6v-2h4v4h-2v-2h-2zm-2 0h-2v4h-2v2h4v-6zm-6 0v-2h4v2zM6 22h4v4H6z"/><path fill="currentColor" d="M14 30H2V18h12zM4 28h8v-8H4zM22 6h4v4h-4z"/><path fill="currentColor" d="M30 14H18V2h12zm-10-2h8V4h-8zM6 6h4v4H6z"/><path fill="currentColor" d="M14 14H2V2h12ZM4 12h8V4H4Z"/>`),
		g.Group(children),
	)
}