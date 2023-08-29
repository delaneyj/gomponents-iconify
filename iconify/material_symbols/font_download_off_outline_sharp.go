package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontDownloadOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-1.3-1.3H2V4.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM4 20h13.175L4 6.825V20Zm18-.875l-2-2V4H6.875l-2-2H22v17.125Zm-6.4-6.4l-3.05-3.05l-.5-1.425h-.1l-.225.6l-1.35-1.35l.575-1.5h2.1l2.55 6.725Zm-5 .675Zm2.85-2.85ZM15.5 18l-1.425-3.95l3.275 3.275l.25.675h-2.1Zm-9.1 0l3.2-8.425l1.4 1.4l-.8 2.225h3.025l1.75 1.75H9.6L8.5 18H6.4Z"/>`),
		g.Group(children),
	)
}