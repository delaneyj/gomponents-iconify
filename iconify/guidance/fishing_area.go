package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishingArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.093 2.093C2.324 1.351 2.5 1 2.5 1H3s1.825 1.5 4 1.5h.035m-2.88 1.656A8.99 8.99 0 0 1 7.036 2.5m1.027 9.062C6.832 12.33 6 13.806 6 15.5c0 2.485 1.79 4.5 4 4.5m0 0A9 9 0 0 1 2.51 6.01M10 20c1.078 1.931 3.093 3.5 5.4 3.5v-.25S13.5 22 13.5 20c0-.672.215-1.26.5-1.744M10 20c.68-1.22 1.735-2.294 3-2.921M21.25 0v6.25l.25.25A1.5 1.5 0 1 1 20 8m-8-2V4M.5.5l10.499 10.499M23.5 23.5L10.999 10.999m0 0c2.41-.027 4.054-.552 6.001-2.499v-.25c-2-.25-3-.75-3-.75v-.25S15.5 6 17.25 5.5v-.25C15.724 3.223 13.401 2 10.8 2H10a8.987 8.987 0 0 0-2.965.5"/>`),
		g.Group(children),
	)
}