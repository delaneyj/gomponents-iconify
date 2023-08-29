package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 172a40 40 0 0 1-80 0V76h-56v124a12 12 0 0 1-24 0V76h-4a36 36 0 0 0-36 36a12 12 0 0 1-24 0a60.07 60.07 0 0 1 60-60h152a12 12 0 0 1 0 24h-44v96a16 16 0 0 0 32 0a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}