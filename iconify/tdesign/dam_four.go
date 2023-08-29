package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h4V10h3V4H4Zm9 0v6h3v10h4V4h-7Zm1 16v-8h-4v8h4Z"/>`),
		g.Group(children),
	)
}