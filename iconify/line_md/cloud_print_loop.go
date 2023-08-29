package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudPrintLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudPrintLoop0"><g fill="#fff"><circle cx="12" cy="8" r="6"/><rect width="15" height="12" x="1" y="6" rx="6"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="8" rx="5"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="12" height="11" x="6" y="11" fill="#fff"/><rect width="8" height="7" x="8" y="13"/><path fill="#fff" d="M9 12h6v1H9zM9 14h6v1H9zM9 16h6v1H9zM9 18h6v1H9z"><animateMotion dur="1.5s" path="M0 0v2" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudPrintLoop0)"/>`),
		g.Group(children),
	)
}