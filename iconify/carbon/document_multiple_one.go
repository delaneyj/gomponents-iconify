package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentMultipleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 6h2v20H2zm4-2h2v24H6zm8 18h12v2H14zm0-6h12v2H14z"/><path fill="currentColor" d="m29.7 9.3l-7-7c-.2-.2-.4-.3-.7-.3H12c-1.1 0-2 .9-2 2v24c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V10c0-.3-.1-.5-.3-.7zM22 4.4l5.6 5.6H22V4.4zM28 28H12V4h8v6c0 1.1.9 2 2 2h6v16z"/>`),
		g.Group(children),
	)
}