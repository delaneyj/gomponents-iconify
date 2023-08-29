package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeAcr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.03 14.23L24 24.03l16.86-9.73l-16.97-9.8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.03 33.7L24 43.5V24.03l-16.97-9.8ZM24 43.5l16.86-9.73V14.3L24 24.03Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.347 27.33c.433-.194 2.06-.92 2.211-.979c.198-.12.26-.404.162-.75q-.201-.989-.31-1.941a1.193 1.193 0 0 0-.624-.858l-.019-.009l-2.374-1.37c-.205-.1-.365-.027-.379.17a14.666 14.666 0 0 0 1.287 5.63h0l.01.023l.036.083h.002a22.557 22.557 0 0 0 3.634 5.723l-.004.001l.122.139l.004.004h0a19.617 19.617 0 0 0 4.837 4.056c.205.104.37.039.389-.155v-2.376a1.099 1.099 0 0 0-.476-.916l-.019-.013q-.888-.621-1.76-1.326a.99.99 0 0 0-.844-.325c-.146.068-1.808.734-2.253.913m17.849-9.457h0c1.064-.28 1.961.538 2.004 1.826h0l.094 2.883a2.861 2.861 0 0 1-1.85 2.84h0c-1.064.28-1.96-.539-2.003-1.827v0l-.095-2.883a2.862 2.862 0 0 1 1.85-2.839Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.765 30.522a2.853 2.853 0 0 0 4.166 2.047a4.736 4.736 0 0 0 2.42-3.806m-2.854 3.987v2.605m-3.625 1.949l7.007-3.767"/>`),
		g.Group(children),
	)
}