package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="35.958" r="28"/><path stroke-linecap="round" d="M45.429 46.458v-22L36 43.315l-9.429-18.857v22"/></g><circle cx="36" cy="36" r="28" fill="#1e50a0"/><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><path stroke-linecap="round" d="M45.429 46.5v-22L36 43.357L26.571 24.5v22"/></g>`),
		g.Group(children),
	)
}