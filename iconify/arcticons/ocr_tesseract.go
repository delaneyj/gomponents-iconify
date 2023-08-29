package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OcrTesseract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.66 22.059h32.68a3.16 3.16 0 0 1 3.16 3.16v10.2a3.16 3.16 0 0 1-3.16 3.16H7.66a3.16 3.16 0 0 1-3.16-3.16v-10.2a3.16 3.16 0 0 1 3.16-3.16Zm32.66.004L5.37 9.421"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.236 28.144h22.338v4.35H17.236zm-8.899 0h4.343v4.35H8.337z"/>`),
		g.Group(children),
	)
}