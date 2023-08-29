package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M445.7 127.9V384l-63.4 36.5V164.7L223.8 73.1L65.2 164.7l.4 255.9L2.3 384V128.1L224.2 0l221.5 127.9zM255.6 420.5L224 438.9l-31.8-18.2v-256l-63.3 36.6l.1 255.9l94.9 54.9l95.1-54.9v-256l-63.4-36.6v255.9z"/>`),
		g.Group(children),
	)
}