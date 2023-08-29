package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstContribution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.716.315a1 1 0 0 0-1.432 0L6.646.97a1 1 0 0 1-.988.265l-.88-.249a1 1 0 0 0-1.24.717l-.226.886a1 1 0 0 1-.723.723l-.886.225a1 1 0 0 0-.717 1.24l.249.881a1 1 0 0 1-.265.987l-.655.639a1 1 0 0 0 0 1.432l.655.639a1 1 0 0 1 .265.987l-.249.88a1 1 0 0 0 .717 1.24l.886.226a1 1 0 0 1 .723.723l.225.886a1 1 0 0 0 1.24.716l.881-.248a1 1 0 0 1 .988.265l.638.655a1 1 0 0 0 1.432 0l.639-.655a1 1 0 0 1 .987-.265l.88.248a1 1 0 0 0 1.24-.716l.226-.886a1 1 0 0 1 .723-.723l.886-.225a1 1 0 0 0 .716-1.24l-.248-.881a1 1 0 0 1 .265-.988l.655-.638a1 1 0 0 0 0-1.432l-.655-.638a1 1 0 0 1-.265-.988l.248-.88a1 1 0 0 0-.716-1.24l-.886-.226a1 1 0 0 1-.723-.723l-.225-.886a1 1 0 0 0-1.24-.717l-.881.249A1 1 0 0 1 9.354.97L8.716.315ZM7 4h2v8H7V6H6V5l1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}