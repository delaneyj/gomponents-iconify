package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M249.89 76.47a4 4 0 0 0-4.11.2L204 104.53V72a12 12 0 0 0-12-12h-78.94a4 4 0 0 0 0 8H192a4 4 0 0 1 4 4v87.63a4 4 0 0 0 8 0v-8.16l41.78 27.86A4 4 0 0 0 252 176V80a4 4 0 0 0-2.11-3.53ZM244 168.53l-40-26.67v-27.72l40-26.67ZM51 37.31a4 4 0 0 0-6 5.38L60.78 60H32a12 12 0 0 0-12 12v112a12 12 0 0 0 12 12h152.41L205 218.69a4 4 0 1 0 5.92-5.38ZM32 188a4 4 0 0 1-4-4V72a4 4 0 0 1 4-4h36.05l109.09 120Z"/>`),
		g.Group(children),
	)
}