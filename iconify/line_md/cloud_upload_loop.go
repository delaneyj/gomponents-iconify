package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUploadLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudUploadLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="4" height="5" x="10" y="12"/><path d="M12 8L17 13H7L12 8Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v-1v2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudUploadLoop0)"/>`),
		g.Group(children),
	)
}