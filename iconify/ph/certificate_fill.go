package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 88.86V56a16 16 0 0 0-16-16H40a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h112v24a8 8 0 0 0 11.58 7.16L192 216.94l28.42 14.22A8 8 0 0 0 232 224v-56.86a55.87 55.87 0 0 0 0-78.28ZM128 144H72a8 8 0 0 1 0-16h56a8 8 0 0 1 0 16Zm0-32H72a8 8 0 0 1 0-16h56a8 8 0 0 1 0 16Zm88 99.06l-20.42-10.22a8 8 0 0 0-7.16 0L168 211.06v-32.47a55.94 55.94 0 0 0 48 0ZM192 168a40 40 0 1 1 40-40a40 40 0 0 1-40 40Z"/>`),
		g.Group(children),
	)
}