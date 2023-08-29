package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocameraRecordBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12.5v-1c0-3.287 0-4.931.908-6.038a4 4 0 0 1 .554-.554C4.57 4 6.212 4 9.5 4c3.287 0 4.931 0 6.038.908a4 4 0 0 1 .554.554c.702.855.861 2.031.897 4.038l.67-.33c1.945-.972 2.918-1.459 3.63-1.019c.711.44.711 1.528.711 3.703v.292c0 2.175 0 3.263-.711 3.703c-.712.44-1.685-.047-3.63-1.02l-.67-.329c-.036 2.007-.195 3.183-.897 4.038a4.001 4.001 0 0 1-.554.554C14.43 20 12.788 20 9.5 20c-3.287 0-4.931 0-6.038-.908a4 4 0 0 1-.554-.554C2 17.43 2 15.788 2 12.5Zm11.56-2.94a1.5 1.5 0 1 0-2.121-2.12a1.5 1.5 0 0 0 2.122 2.12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}