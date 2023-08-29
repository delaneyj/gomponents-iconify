package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesignServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.85 9.4L14.6 5.15l1.425-1.425q.575-.575 1.413-.575t1.412.575l1.425 1.425q.575.575.575 1.413t-.575 1.412L18.85 9.4ZM3 21v-4.25l4.5-4.5l-5.275-5.325L6.95 2.2l5.3 5.325l.95-.95l4.25 4.225l-.95.95l5.275 5.325l-4.7 4.7l-5.325-5.3L7.25 21H3Zm5.925-10.175l1.9-1.9l-1.2-1.2l-1.2 1.175l-1.4-1.4L8.2 6.3L6.925 5.05L5.05 6.95l3.875 3.875Zm8.125 8.125l1.9-1.9l-1.275-1.25l-1.175 1.175l-1.425-1.4l1.2-1.2l-1.225-1.2l-1.9 1.9l3.9 3.875Z"/>`),
		g.Group(children),
	)
}