package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTopbuzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.417 8.655a.524.524 0 0 1-.405-.622l.986-4.617a.524.524 0 0 1 .626-.404l14.958 3.162c.285.06.467.339.406.622l-.987 4.618a.524.524 0 0 1-.625.404l-4.345-.92c-.198-.04-.315.024-.353.197l-2.028 9.49a.527.527 0 0 1-.625.404l-4.642-.982a.527.527 0 0 1-.406-.622l2.028-9.493c.037-.17-.031-.274-.204-.31l-4.384-.927z"/>`),
		g.Group(children),
	)
}