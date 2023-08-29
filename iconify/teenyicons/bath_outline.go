package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 7.5h15m-10.5 5h6m-6 0a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h2V2m-2 10.5V15m6-2.5a3 3 0 0 0 3-3v-2m-3 5V15M5 3.5h3"/>`),
		g.Group(children),
	)
}