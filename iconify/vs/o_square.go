package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm221 896q0-142 99.5-238.5T899 560q139 0 237.5 98.5T1235 896q0 142-99.5 239T893 1232q-139 0-237.5-98.5T557 896zm339 597q162 0 299.5-80t217.5-217.5t80-299.5t-80-299.5T1195.5 379T896 299t-299.5 80T379 596.5T299 896t80 299.5T596.5 1413t299.5 80z"/>`),
		g.Group(children),
	)
}