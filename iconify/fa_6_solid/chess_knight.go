package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessKnight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 48L82.7 61.3c-12 12-18.7 28.2-18.7 45.2v132.4c0 10.7 5.3 20.7 14.2 26.6l10.6 7c14.3 9.6 32.7 10.7 48.1 3l3.2-1.6c2.6-1.3 5-2.8 7.3-4.5l49.4-37c6.6-5 15.7-5 22.3 0c10.2 7.7 9.9 23.1-.7 30.3L90.4 350C73.9 361.3 64 380 64 400h320l28.9-159c2.1-11.3 3.1-22.8 3.1-34.3V192C416 86 330 0 224 0H83.8C72.9 0 64 8.9 64 19.8c0 7.5 4.2 14.3 10.9 17.7L96 48zm24 68a20 20 0 1 1 40 0a20 20 0 1 1-40 0zM22.6 473.4c-4.2 4.2-6.6 10-6.6 16c0 12.5 10.1 22.6 22.6 22.6h370.8c12.5 0 22.6-10.1 22.6-22.6c0-6-2.4-11.8-6.6-16L384 432H64l-41.4 41.4z"/>`),
		g.Group(children),
	)
}