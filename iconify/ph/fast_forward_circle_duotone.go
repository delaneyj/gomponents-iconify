package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 32a96 96 0 1 0 96 96a96 96 0 0 0-96-96ZM80 160V96l48 32Zm64 0V96l48 32Z" opacity=".2"/><path d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm4.44-94.66l-48-32A8 8 0 0 0 72 96v64a8 8 0 0 0 12.44 6.66l48-32a8 8 0 0 0 0-13.32ZM88 145.05V111l25.58 17Zm108.44-23.71l-48-32A8 8 0 0 0 136 96v64a8 8 0 0 0 12.44 6.66l48-32a8 8 0 0 0 0-13.32ZM152 145.05V111l25.58 17Z"/></g>`),
		g.Group(children),
	)
}