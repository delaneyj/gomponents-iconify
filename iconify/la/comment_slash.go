package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.707 2.293L2.293 3.707l26 26l1.414-1.414L27.414 26H29V6H7.414L3.707 2.293zM3 7.243V26h9.586L16 29.414L19.414 26h2.344l-2-2h-1.172L16 26.586L13.414 24H5V9.242l-2-2zM9.414 8H27v16h-1.586l-16-16z"/>`),
		g.Group(children),
	)
}