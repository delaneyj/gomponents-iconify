package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUploadOutlineLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdCloudUploadOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"/><rect width="8" height="8" x="8" y="10"/><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="#fff"><rect width="3" height="4" x="10.5" y="12"/><path d="M12 9L16 13H8L12 9Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v-1v2z" repeatCount="indefinite"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudUploadOutlineLoop0)"/>`),
		g.Group(children),
	)
}