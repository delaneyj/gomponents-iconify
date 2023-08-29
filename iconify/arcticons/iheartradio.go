package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iheartradio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 40.56V26.523"/><circle cx="24" cy="20.084" r="3.019" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.401 17.056A10.536 10.536 0 0 0 24 12.976A10.535 10.535 0 1 0 8.645 26.881h-.018l13.978 13.136a2 2 0 0 0 2.739 0l13.848-13.014a10.522 10.522 0 0 0 4.21-9.948Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.322 23.734a7.309 7.309 0 0 0 0-7.3m3.668 9.418a11.547 11.547 0 0 0 0-11.535m-16.312 2.117a7.309 7.309 0 0 0 0 7.3m-3.668-9.417a11.547 11.547 0 0 0 0 11.535"/>`),
		g.Group(children),
	)
}