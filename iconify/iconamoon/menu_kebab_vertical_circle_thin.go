package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuKebabVerticalCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`),
		g.Group(children),
	)
}