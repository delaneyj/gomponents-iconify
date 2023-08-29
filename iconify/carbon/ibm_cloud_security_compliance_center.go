package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudSecurityComplianceCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 18v5.5c0 1.5-.9 2.8-2.3 3.5l-1.7.8l-1.7-.8c-1.4-.6-2.3-2-2.3-3.5V18h8m2-2H20v7.5c0 2.2 1.3 4.3 3.4 5.3L26 30l2.6-1.2c2.1-1 3.4-3 3.4-5.3V16z"/><path fill="currentColor" d="M16 25H7.5C3.4 25 0 21.6 0 17.5c0-3.7 2.7-6.7 6.2-7.4c.9-4.7 5-8.1 9.8-8.1c5.5 0 10 4.5 10 10h-2c0-4.4-3.6-8-8-8c-4.1 0-7.5 3.1-8 7.1v.9h-.9c-2.9.2-5.1 2.6-5.1 5.5c0 3 2.5 5.5 5.5 5.5H16v2z"/>`),
		g.Group(children),
	)
}