package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EslintOld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M303.26 0L54.492 80.844L0 336.885l194.525 175.114l248.768-80.844l54.492-256.041L303.26 0zM181.1 138.537h135.585L384.561 256l-67.876 117.463H181.1L113.432 256L181.1 138.537z"/>`),
		g.Group(children),
	)
}