package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sonic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v20h-2V2h2ZM9 6v12H7V6h2Zm8 0v12h-2V6h2ZM5 9v6H3V9h2Zm16 0v6h-2V9h2Z"/>`),
		g.Group(children),
	)
}