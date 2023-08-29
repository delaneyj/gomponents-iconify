package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConceptSharing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConceptSharing0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M28.367 36H24l-9-4.96l-10.991 4.042l3.002 5.944l7.072-2.953C20.7 42.024 24.727 44 26.165 44c1.438 0 7.383-2.667 17.835-8"/><path fill="#fff" fill-rule="evenodd" stroke-linejoin="round" d="M28.992 26.988v-4.67c1.1-.457 2.543-1.125 3.372-1.954a9 9 0 1 0-12.728 0c.829.829 2.264 1.497 3.364 1.953c.006.335.006 1.892 0 4.67h5.992Z" clip-rule="evenodd"/><path stroke-linecap="round" d="m12 21l1-1m27 1l-1-1M15 5l-1-1m23 1l1-1m3 8h-1m-28 0h-1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConceptSharing0)"/>`),
		g.Group(children),
	)
}