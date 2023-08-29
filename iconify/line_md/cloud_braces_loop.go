package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudBracesLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudBracesLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="10;11;10;9;10"/></rect></g><path d="M18.5 12H18a1 1 0 0 1-1-1v-1a2 2 0 0 0-2-2h-1.5v2H15v1a2 2 0 0 0 2 2 2 2 0 0 0-2 2v1h-1.5v2H15a2 2 0 0 0 2-2v-1a1 1 0 0 1 1-1h.5v-2Z"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h-1h2z" repeatCount="indefinite"/></path><path d="M5.5 12v2H6a1 1 0 0 1 1 1v1a2 2 0 0 0 2 2h1.5v-2H9v-1a2 2 0 0 0-2-2 2 2 0 0 0 2-2v-1h1.5V8H9a2 2 0 0 0-2 2v1a1 1 0 0 1-1 1h-.5Z"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h1h-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudBracesLoop0)"/>`),
		g.Group(children),
	)
}