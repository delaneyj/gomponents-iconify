package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xmake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m8.12 288.788l178.604-110.14l288.43 177.865C366.642 538.787 73.653 522.044 8.118 288.788zM424.671 31.914L205.773 166.9l279.045 172.08c54.574-115.763 23.146-232.327-60.146-307.065zM3.664 268.04l164.01-101.14L39.385 87.79C6.094 140.99-7.08 200.408 3.665 268.04z"/>`),
		g.Group(children),
	)
}