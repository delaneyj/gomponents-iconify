package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrganizationThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 2a5 5 0 0 0-1.001 9.9v3.099H9.733A2.733 2.733 0 0 0 7 17.732V20.1A5.002 5.002 0 0 0 8 30a5 5 0 0 0 1-9.9v-2.368c0-.405.329-.733.733-.733h12.534c.405 0 .733.328.733.733V20.1a5.002 5.002 0 0 0 1 9.9a5 5 0 0 0 1-9.9v-2.368a2.733 2.733 0 0 0-2.733-2.733H17V11.9A5.002 5.002 0 0 0 16 2Zm-3 5a3 3 0 1 1 6 0a3 3 0 0 1-6 0ZM5 25a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm19-3a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z"/>`),
		g.Group(children),
	)
}