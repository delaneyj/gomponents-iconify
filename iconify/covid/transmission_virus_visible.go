package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionVirusVisible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.824 20.36a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858Zm-.572-9.429h1.143m-.571 0v2.571m3.838-1.218l.808.808m-.404-.404l-1.818 1.818m3.576 1.853v1.143m0-.571h-2.572m1.218 3.838l-.808.808m.404-.404l-1.818-1.818m-1.853 3.576h-1.143m.572 0v-2.572m-3.839 1.218l-.808-.808m.404.404l1.818-1.818m-3.575-1.853v-1.143m0 .572h2.571m-1.218-3.839l.808-.808m-.404.404l1.818 1.818m7.648-4.889c.254-.253.498-.508.727-.762a1.667 1.667 0 0 0 0-2.226C20.163 3.746 16.012.924 11.998.995C7.986.925 3.835 3.746 1.226 6.63a1.667 1.667 0 0 0 0 2.226a19.308 19.308 0 0 0 6.287 4.647"/><path d="M9.256 10.376a3.751 3.751 0 1 1 6.466-2.091"/></g>`),
		g.Group(children),
	)
}