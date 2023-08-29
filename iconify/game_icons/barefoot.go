package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barefoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M44.156 19.47c40.24 103.666 45.06 254.208 5.22 367.28c-28.324 80.378 9.553 105.938 66.25 105.938l31.374.28c63.568-40.225 127.15-23.413 190.72 1.813l110.78 1.033c77.705 0 54.114-80.826-18.594-69.125L278.53 349.5C212.117 277.39 222.04 85.354 241.407 20l-197.25-.53z"/>`),
		g.Group(children),
	)
}