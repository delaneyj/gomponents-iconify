package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudPrintOutlineLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudPrintOutlineLoop0"><g fill="#fff"><circle cx="12" cy="8" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="17" height="12" x="1" y="6" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="8" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="8" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="8"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="8" rx="4"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="10" rx="3"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="12" height="11" x="6" y="11" fill="#fff"/><rect width="8" height="7" x="8" y="13"/><path fill="#fff" d="M9 12h6v1H9zM9 14h6v1H9zM9 16h6v1H9zM9 18h6v1H9z"><animateMotion dur="1.5s" path="M0 0v2" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudPrintOutlineLoop0)"/>`),
		g.Group(children),
	)
}