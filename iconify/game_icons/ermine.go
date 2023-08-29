package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ermine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 24l-32 64l32 48l32-48l-32-64zm-64 96l-64 32l64 32l48-32l-48-32zm128 0l-48 32l48 32l64-32l-64-32zm-64 32c-32 128-64 192-128 256c16 0 48 0 64-16c0 16-16 48-32 64c16 0 48 0 64-16c16 16 16 16 32 48c16-32 16-32 32-48c16.847 12.064 48 16 64 16c-16-16-32-48-32-64c16 16 48 16 64 16c-64-64-96-128-128-256z"/>`),
		g.Group(children),
	)
}