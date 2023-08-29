package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M35 16h25v40H35z"/><path fill="#D0CFCE" d="M60.053 56V28L35 56"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M35 16h25v40H35zm-5 20H11m13-6l6 6m-6 6l6-6"/>`),
		g.Group(children),
	)
}