package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighHeelLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230.53 166.14L179 154.53L68.24 43.76a6 6 0 0 0-8.67.19C37.92 67.62 26 97.47 26 128v72a14 14 0 0 0 14 14h32a14 14 0 0 0 14-14v-52.35a115 115 0 0 1 51.74 57.69a13.92 13.92 0 0 0 13 8.66H240a14 14 0 0 0 14-14v-4.73a29.73 29.73 0 0 0-23.47-29.13ZM74 200a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2v-65.82A112.75 112.75 0 0 1 74 142Zm168 0a2 2 0 0 1-2 2h-89.28a2 2 0 0 1-1.86-1.18a126.53 126.53 0 0 0-45.58-56.65a125.13 125.13 0 0 0-65.12-22c1.31-23.51 10.44-46.34 26.15-65.38l107.45 107.45a6 6 0 0 0 2.92 1.61l53.23 12A17.81 17.81 0 0 1 242 195.31Z"/>`),
		g.Group(children),
	)
}