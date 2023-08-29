package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Six(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSix0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 41C13.5 38.25 8 30 5 25s3.313-9.688 7-6l4 4v-5.5a3.5 3.5 0 1 1 7 0V16a3.5 3.5 0 1 1 7 0v1.5a3.5 3.5 0 1 1 7 0v-10a3.5 3.5 0 1 1 7 0v21.466c0 2.003-.37 4.008-1.456 5.691C41.264 36.645 39.111 39.303 36 41c-5.5 3-11.5 2.75-17 0Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSix0)"/>`),
		g.Group(children),
	)
}