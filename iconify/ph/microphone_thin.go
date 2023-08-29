package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 172a44.05 44.05 0 0 0 44-44V64a44 44 0 0 0-88 0v64a44.05 44.05 0 0 0 44 44ZM92 64a36 36 0 0 1 72 0v64a36 36 0 0 1-72 0Zm40 139.89V232a4 4 0 0 1-8 0v-28.11A76.09 76.09 0 0 1 52 128a4 4 0 0 1 8 0a68 68 0 0 0 136 0a4 4 0 0 1 8 0a76.09 76.09 0 0 1-72 75.89Z"/>`),
		g.Group(children),
	)
}