package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailCom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.925 18.103a6.038 6.038 0 0 1 6.037-6.038h0A6.038 6.038 0 0 1 24 18.103v9.962m-12.075-16v16M24 18.103a6.038 6.038 0 0 1 6.038-6.038h0a6.038 6.038 0 0 1 6.037 6.038v9.962"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.494 42L8.087 34.85a2 2 0 0 1-1.581-1.956V6h31.988a3 3 0 0 1 3 3v33Z"/>`),
		g.Group(children),
	)
}