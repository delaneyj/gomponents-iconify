package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.172 4.244c-.457.541-4.619 4.77-6.846 6.422l-.913-.867c1.29-2.511 4.741-7.292 5.196-7.832c.595-.704 2.037-1.215 2.807-.669c1.012.716.35 2.24-.244 2.946zm-8.565 6.458s-2.092-.553-3.348 1.594C2.004 14.443.808 13.708.27 13.524c-.538-.184 1.256 2.27 2.093 2.27c.818 0 4.389 1.175 5.268-3.434c.018-.101.072-1.483-1.024-1.658z"/>`),
		g.Group(children),
	)
}