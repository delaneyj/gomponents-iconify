package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Click(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClick0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 4v8"/><path fill="#fff" fill-rule="evenodd" d="m22 22l20 4l-6 4l6 6l-6 6l-6-6l-4 6l-4-20Z" clip-rule="evenodd"/><path d="m38.142 9.858l-5.657 5.657M9.858 38.142l5.657-5.657M4 24h8M9.858 9.858l5.657 5.657"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClick0)"/>`),
		g.Group(children),
	)
}