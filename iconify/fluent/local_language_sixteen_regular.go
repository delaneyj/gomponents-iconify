package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLanguageSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 1.996a.5.5 0 0 1 .5.5v2.506h.5a.5.5 0 0 1 0 1H13V10.5a.5.5 0 1 1-1 0V2.496a.5.5 0 0 1 .5-.5Zm-5 1H10V4.25c0 .413-.158.708-.4.908c-.251.209-.63.342-1.097.342a.5.5 0 0 0 0 1c.648 0 1.268-.184 1.735-.572c.479-.396.762-.976.762-1.678V2.496a.5.5 0 0 0-.5-.5h-3a.5.5 0 1 0 0 1ZM6.454 5.3a.5.5 0 0 0-.916 0l-3.496 8a.5.5 0 1 0 .916.401l1.18-2.7H7.86l1.183 2.7a.5.5 0 0 0 .916-.4l-3.504-8Zm.967 4.7H4.575l1.422-3.252L7.42 10Z"/>`),
		g.Group(children),
	)
}