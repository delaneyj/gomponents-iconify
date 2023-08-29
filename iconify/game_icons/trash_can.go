package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashCan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M199 103v50h-78v30h270v-30h-78v-50H199zm18 18h78v32h-78v-32zm-79.002 80l30.106 286h175.794l30.104-286H137.998zm62.338 13.38l.64 8.98l16 224l.643 8.976l-17.956 1.283l-.64-8.98l-16-224l-.643-8.976l17.956-1.283zm111.328 0l17.955 1.284l-.643 8.977l-16 224l-.64 8.98l-17.956-1.284l.643-8.977l16-224l.64-8.98zM247 215h18v242h-18V215z"/>`),
		g.Group(children),
	)
}