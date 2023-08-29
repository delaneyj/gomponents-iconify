package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TidalLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m252.49 87.51l-38-38a12 12 0 0 0-17 0L168 79l-31.51-31.49a12 12 0 0 0-17 0L88 79L58.49 49.51a12 12 0 0 0-17 0l-38 38a12 12 0 0 0 0 17l38 38a12 12 0 0 0 17 0L88 113l23 23l-31.49 31.51a12 12 0 0 0 0 17l40 40a12 12 0 0 0 17 0l40-40a12 12 0 0 0 0-17L145 136l23-23l29.51 29.52a12 12 0 0 0 17 0l38-38a12 12 0 0 0-.02-17.01ZM50 117L29 96l21-21l21 21Zm78 82l-23-23l23-23l23 23Zm0-80l-23-23l23-23l23 23Zm78-2l-21-21l21-21l21 21Z"/>`),
		g.Group(children),
	)
}