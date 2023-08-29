package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m341 21l86 86h-64v149q0 35-25 60t-60.5 25t-60.5-25t-25-60V107q0-18-12.5-30.5t-30-12.5t-30 12.5T107 107v149h64l-86 85l-85-85h64V107q0-36 25-61t60.5-25T210 46t25 61v149q0 18 12.5 30.5t30 12.5t30-12.5T320 256V107h-64z"/>`),
		g.Group(children),
	)
}