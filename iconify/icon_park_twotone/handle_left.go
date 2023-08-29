package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandleLeft0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linecap="round" rx="3"/><path d="M28 34L18 24l10-10v20Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandleLeft0)"/>`),
		g.Group(children),
	)
}