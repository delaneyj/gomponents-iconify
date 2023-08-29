package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="22" fill="#3F3F3F"/><circle cx="36" cy="36" r="12.756" fill="#9B9B9A"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="22"/><circle cx="36" cy="36" r="12.756"/></g>`),
		g.Group(children),
	)
}