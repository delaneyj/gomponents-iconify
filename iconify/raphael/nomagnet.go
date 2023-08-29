package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nomagnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.59 17.857v-5.225c-.027-.376.303-1.79 1.1-2.748c.818-.98 1.848-1.748 4.013-1.778c1.704.026 2.7.508 3.447 1.19l3.54-3.54c-1.617-1.526-4.01-2.68-6.987-2.652c-3.626-.03-6.403 1.675-7.94 3.69c-1.563 2.032-2.146 4.134-2.173 5.84V19.5h3.357l1.643-1.643zm-5 2.643v2.357L7.947 20.5H5.59zm15.222-7.21v6.21h5.002v-6.866a9.314 9.314 0 0 0-.803-3.542l-4.198 4.198zm4.527-8.768L4.65 25.21l1.415 1.415L26.753 5.937L25.34 4.522zM20.81 25.58h5.002V20.5H20.81v5.08zm-10.222 0v-2.064L8.525 25.58h2.065z"/>`),
		g.Group(children),
	)
}