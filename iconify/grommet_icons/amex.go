package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4.314 11.965l-.82-1.997l-.815 1.997h1.635Zm7.859 2.161l-.005-3.922l-1.736 3.922h-1.05L7.64 10.2v3.926H5.206l-.46-1.117H2.253l-.465 1.117h-1.3l2.144-5.008H4.41l2.036 4.742V9.118H8.4l1.567 3.397l1.439-3.397H13.4v5.008h-1.227Zm3.133-1.024v-.997h2.628v-1.022h-2.628v-.911h3.001l1.31 1.46l-1.368 1.47h-2.943Zm8.111 1.044h-1.556l-1.474-1.659l-1.532 1.659h-4.742v-5.01h4.815l1.473 1.642l1.523-1.642h1.564l-2.327 2.505l2.256 2.505Z"/>`),
		g.Group(children),
	)
}