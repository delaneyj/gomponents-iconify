package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M51 37.31a4 4 0 0 0-6 5.38l1.31 1.44A12 12 0 0 0 36 56v58.77c0 86.88 73.54 115.7 88.28 120.59a11.47 11.47 0 0 0 7.44 0c7.82-2.59 34.16-12.64 55.95-35.8L205 218.69a4 4 0 1 0 5.92-5.38Zm78.23 190.48a3.51 3.51 0 0 1-2.39 0C113 223.2 44 196.17 44 114.79V56a4 4 0 0 1 4-4h5.5l128.77 141.65c-20.58 22.12-45.66 31.68-53.08 34.14ZM220 56v58.77c0 19.1-3.62 36.81-10.75 52.65a4 4 0 0 1-7.29-3.29c6.66-14.79 10-31.4 10-49.36V56a4 4 0 0 0-4-4H98.52a4 4 0 1 1 0-8H208a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}