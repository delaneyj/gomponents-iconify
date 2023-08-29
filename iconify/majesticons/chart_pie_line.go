package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 12c0-3.732 2.554-6.907 6-7.802V11a3 3 0 0 0 3 3h6.802c-.895 3.446-4.07 6-7.802 6a8 8 0 0 1-8-8zm8-7.883c0-1.15-1.037-2.21-2.334-1.897C5.264 3.28 2 7.285 2 12c0 5.523 4.477 10 10 10c4.715 0 8.72-3.264 9.78-7.666c.312-1.296-.748-2.334-1.897-2.334H13a1 1 0 0 1-1-1V4.117zM15 3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 .98-1.196c-.19-.952-.693-2.354-1.615-3.542C18.429 4.055 16.997 3 15 3zm1 5V5.155c.722.234 1.308.718 1.785 1.333c.37.476.654 1.01.862 1.512H16z"/></g>`),
		g.Group(children),
	)
}