package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HollowRedCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#EA5A47" d="M35.566 6.407c-16.016 0-29 12.983-29 29c0 16.016 12.984 29 29 29s29-12.984 29-29c0-16.017-12.984-29-29-29zm0 49c-11.046 0-20-8.955-20-20s8.954-20 20-20s20 8.954 20 20s-8.954 20-20 20z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="35.795" cy="35.937" r="29"/><circle cx="35.795" cy="35.937" r="20"/></g>`),
		g.Group(children),
	)
}