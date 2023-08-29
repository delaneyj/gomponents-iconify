package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 44H32a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h192a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20Zm-4 144H36V68h184ZM48 128a12 12 0 0 1 12-12h16a12 12 0 0 1 0 24H60a12 12 0 0 1-12-12Zm56 0a12 12 0 0 1 12-12h80a12 12 0 0 1 0 24h-80a12 12 0 0 1-12-12Zm-56 36a12 12 0 0 1 12-12h80a12 12 0 0 1 0 24H60a12 12 0 0 1-12-12Zm160 0a12 12 0 0 1-12 12h-16a12 12 0 0 1 0-24h16a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}