package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 8v11c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.806 5 19s.55-.988 1.012-1H21V4c0-1.103-.897-2-2-2H6c-1.206 0-3 .799-3 3v3zm3-4h13v12H5V5c0-.806.55-.988 1-1z"/><path fill="currentColor" d="m11.997 14l3.35-3.289a2.129 2.129 0 0 0 0-3.069a2.225 2.225 0 0 0-3.126 0l-.224.218l-.224-.219a2.224 2.224 0 0 0-3.125 0a2.129 2.129 0 0 0 0 3.069L11.997 14z"/>`),
		g.Group(children),
	)
}