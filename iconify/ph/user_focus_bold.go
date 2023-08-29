package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFocusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 48v28a12 12 0 0 1-24 0V52h-24a12 12 0 0 1 0-24h28a20 20 0 0 1 20 20Zm-12 120a12 12 0 0 0-12 12v24h-24a12 12 0 0 0 0 24h28a20 20 0 0 0 20-20v-28a12 12 0 0 0-12-12ZM76 204H52v-24a12 12 0 0 0-24 0v28a20 20 0 0 0 20 20h28a12 12 0 0 0 0-24ZM40 88a12 12 0 0 0 12-12V52h24a12 12 0 0 0 0-24H48a20 20 0 0 0-20 20v28a12 12 0 0 0 12 12Zm136 92a12 12 0 0 1-9.6-4.79a48 48 0 0 0-76.82 0a12 12 0 0 1-19.18-14.42a72.1 72.1 0 0 1 23.92-20.5a44 44 0 1 1 67.34 0a72.1 72.1 0 0 1 23.92 20.5A12 12 0 0 1 176 180Zm-48-48a20 20 0 1 0-20-20a20 20 0 0 0 20 20Z"/>`),
		g.Group(children),
	)
}