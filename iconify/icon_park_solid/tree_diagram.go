package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeDiagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTreeDiagram0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="10" cy="24" r="4" fill="#fff"/><circle cx="38" cy="10" r="4" fill="#fff"/><circle cx="38" cy="24" r="4" fill="#fff"/><circle cx="38" cy="38" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 38H22V10h12M14 24h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTreeDiagram0)"/>`),
		g.Group(children),
	)
}