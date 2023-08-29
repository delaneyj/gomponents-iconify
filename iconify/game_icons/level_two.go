package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c-22.5 72-45 72-45 120c0 0 22.5 24 45 24s45-24 45-24c0-48-22.5-48-45-120zm2.625 144.03A90 96 0 0 0 166 256a90 96 0 0 0 180 0a90 96 0 0 0-87.375-95.97zM256 352c-22.5 0-45 24-45 24c0 48 22.5 48 45 120c22.5-72 45-72 45-120c0 0-22.5-24-45-24z"/>`),
		g.Group(children),
	)
}