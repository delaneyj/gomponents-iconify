package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yetanothercallblocker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.07 10.1L26.28 4.92a5.61 5.61 0 0 0-4.26 0L8.89 10.1A2 2 0 0 0 7.62 12v12.47c0 11.16 11.6 19 16.58 19s16.18-7.87 16.18-19V12a2 2 0 0 0-1.31-1.9Zm-4.36 23.28a1.37 1.37 0 0 1-1.37 1.35h-.07a21.74 21.74 0 0 1-20-20a1.34 1.34 0 0 1 1.27-1.43h3a2.69 2.69 0 0 1 2.63 2.3a13.31 13.31 0 0 0 .51 2.14a1.58 1.58 0 0 1-.4 1.56L18.62 21A19.22 19.22 0 0 0 27 29.34l1.77-1.75a1.32 1.32 0 0 1 1.38-.33a14.73 14.73 0 0 0 2.25.54a2.8 2.8 0 0 1 2.32 2.66Z"/>`),
		g.Group(children),
	)
}