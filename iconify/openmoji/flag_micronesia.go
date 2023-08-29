package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMicronesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M5 17h62v38H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m41.523 34.356l4.954 1.689l-4.878 1.456l2.939-4.045l.122 5.088l-3.137-4.188zm-11.046 3.288l-4.954-1.689l4.878-1.456l-2.939 4.045l-.122-5.088l3.137 4.188zm3.879-7.167l1.689-4.954l1.456 4.878l-4.045-2.939l5.088-.122l-4.188 3.137zm3.288 11.046l-1.689 4.954l-1.456-4.878l4.045 2.939l-5.088.122l4.188-3.137z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}