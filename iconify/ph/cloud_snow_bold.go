package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 216a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm-64-16a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm-48 0a16 16 0 1 0 16 16a16 16 0 0 0-16-16ZM236 92a80.09 80.09 0 0 1-80 80H76a56 56 0 0 1 0-112a56.89 56.89 0 0 1 6.39.36A80.08 80.08 0 0 1 236 92Zm-24 0a56.06 56.06 0 0 0-112-3.31a12 12 0 1 1-24-1.38c.06-1.11.15-2.21.26-3.31H76a32 32 0 0 0 0 64h80a56.06 56.06 0 0 0 56-56Z"/>`),
		g.Group(children),
	)
}