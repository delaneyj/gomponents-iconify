package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSendBackward0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M6 14h36V6H6zm0 14h36v-8H6z"/><path stroke-linecap="round" d="m30 36l-6 6l-6-6v0m6 6V28m0-14v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSendBackward0)"/>`),
		g.Group(children),
	)
}