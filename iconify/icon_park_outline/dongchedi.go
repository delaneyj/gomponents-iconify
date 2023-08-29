package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dongchedi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M38.16 33.56c-4.316-1.52-8.95-2.315-13.718-2.315c-5.06 0-9.967.895-14.502 2.6a15.53 15.53 0 0 0 1.557 3.142L6.455 40.24A21.69 21.69 0 0 1 3 28.45C3 16.619 12.385 7 24 7s21 9.62 21 21.449a21.68 21.68 0 0 1-3.753 12.24l-4.959-3.378a15.548 15.548 0 0 0 1.873-3.751Zm.44-8.676C37.034 18.061 31.08 13 24 13c-7.172 0-13.188 5.192-14.657 12.15c-.237 1.33.123 2.486 1.11 2.181a47.457 47.457 0 0 1 13.989-2.086c4.657 0 9.212.665 13.532 1.949c.854.253.981-.817.641-2.226a.82.82 0 0 1-.016-.084Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}