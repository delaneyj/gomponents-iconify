package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<circle cx="25" cy="10" r="2" fill="currentColor"/><circle cx="25" cy="40" r="2" fill="currentColor" opacity=".3"/><circle cx="32.5" cy="12" r="2" fill="currentColor" opacity=".3"/><circle cx="17.5" cy="38" r="2" fill="currentColor" opacity=".3"/><circle cx="17.5" cy="12" r="2" fill="currentColor" opacity=".93"/><circle cx="32.5" cy="38" r="2" fill="currentColor" opacity=".3"/><circle cx="10" cy="25" r="2" fill="currentColor" opacity=".65"/><circle cx="40" cy="25" r="2" fill="currentColor" opacity=".3"/><circle cx="12" cy="17.5" r="2" fill="currentColor" opacity=".86"/><circle cx="38" cy="32.5" r="2" fill="currentColor" opacity=".3"/><circle cx="12" cy="32.5" r="2" fill="currentColor" opacity=".44"/><circle cx="38" cy="17.5" r="2" fill="currentColor" opacity=".3"/>`),
		g.Group(children),
	)
}