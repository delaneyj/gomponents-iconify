package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudLoop0)"/>`),
		g.Group(children),
	)
}