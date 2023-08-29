package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#99AAB5" d="M24 10c0-4.971-2.91-10-6.5-10S11 5.029 11 10c0 3.744 1.651 6.385 4 7.461V33.5a2.5 2.5 0 1 0 5 0V17.461c2.349-1.076 4-3.717 4-7.461z"/>`),
		g.Group(children),
	)
}