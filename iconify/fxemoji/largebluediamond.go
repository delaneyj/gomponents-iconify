package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Largebluediamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00B1FF" d="M500.426 271.4c8.378-8.378 8.378-22.088 0-30.466L269.543 10.051c-7.54-7.54-19.879-7.54-27.419 0L9.718 242.457c-7.54 7.54-7.54 19.879 0 27.419L240.6 500.759c8.378 8.378 22.088 8.378 30.466 0L500.426 271.4z"/>`),
		g.Group(children),
	)
}