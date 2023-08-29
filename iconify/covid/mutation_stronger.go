package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MutationStronger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.408.966H9.375m1.517 0v2.6M5.07 2.749L3.997 3.822L2.925 4.894m1.072-1.072L5.836 5.66M1.142 9.2v3.033m0-1.517h2.6m-.817 5.822l1.072 1.073l1.073 1.072m-1.073-1.072l1.839-1.839M18.858 4.894l-1.072-1.072l-1.073-1.073m1.073 1.073L15.947 5.66m-6.572 4.407a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.748 8.683a7.15 7.15 0 1 0-8.89 8.884m3 5.467h8.458a2.533 2.533 0 0 0 2.542-2.523v-6.606a2.189 2.189 0 0 0-1.9-2.211a2.1 2.1 0 0 0-1.152 3.969a5.994 5.994 0 0 0 0 4.428a4.674 4.674 0 0 0-7.338-1.265"/><path d="M11.633 13.625a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		g.Group(children),
	)
}