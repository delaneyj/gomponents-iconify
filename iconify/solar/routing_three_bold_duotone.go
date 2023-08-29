package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoutingThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="5" cy="5" r="3" opacity=".5"/><circle cx="19" cy="19" r="3" opacity=".5"/><path fill-rule="evenodd" d="M10.25 5a.75.75 0 0 1 .75-.75h5.132c2.751 0 3.797 3.593 1.476 5.07l-10.41 6.625c-1.056.672-.58 2.305.67 2.305h3.321l-.22-.22a.75.75 0 1 1 1.061-1.06l1.5 1.5a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l.22-.22H7.867c-2.751 0-3.797-3.593-1.476-5.07l10.411-6.625c1.055-.672.58-2.305-.671-2.305H11a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}