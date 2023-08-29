package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M420 390q26-26 26-61t-26-60l-45-45q-25-26-60-26t-61 26l-30-30q26-26 26-61t-26-60l-45-45Q154 2 119 2T58 28L28 58Q2 84 2 119t26 60l45 45q26 26 59 26q34 0 60-26l30 30q-26 26-26 61t26 60l45 45q26 26 59 26q34 0 60-26zM102 194l-44-45q-13-13-13-29.5T58 90l29-30q12-15 32-15q18 0 30 13l45 44q13 13 13 30t-13 30l-45-43q-14-14-30 0q-14 16 0 30l45 45q-13 13-31 13t-31-13zm197 196l-45-44q-13-13-13-30t13-30l45 45q14 14 30 0q14-16 0-30l-45-45q12-13 30-13q17 0 29 13l45 45q13 13 13 29.5T388 361l-30 29q-11 13-28.5 13T299 390z"/>`),
		g.Group(children),
	)
}