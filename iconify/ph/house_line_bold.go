package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseLineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 204h-12v-88.45a20.07 20.07 0 0 0-6.44-14.7l-79.95-75.47l-.16-.15a19.93 19.93 0 0 0-26.91 0l-.17.15l-79.93 75.47a20.07 20.07 0 0 0-6.44 14.7V204H16a12 12 0 0 0 0 24h224a12 12 0 0 0 0-24ZM52 117.28l76-71.75l76 71.75V204h-40v-44a20 20 0 0 0-20-20h-32a20 20 0 0 0-20 20v44H52ZM140 204h-24v-40h24Z"/>`),
		g.Group(children),
	)
}