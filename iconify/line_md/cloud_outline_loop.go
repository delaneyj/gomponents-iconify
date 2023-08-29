package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOutlineLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="10"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="8;7;8;9;8"/></rect></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudOutlineLoop0)"/>`),
		g.Group(children),
	)
}