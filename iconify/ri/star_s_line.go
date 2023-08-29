package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarSLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17l-5.877 3.59l1.598-6.7l-5.23-4.48l6.865-.55L12 2.5l2.645 6.36l6.865.55l-5.23 4.48l1.598 6.7L12 17Zm0-2.344l2.818 1.72l-.766-3.21l2.506-2.147l-3.29-.264l-1.267-3.047l-1.268 3.047l-3.29.264l2.507 2.147l-.766 3.21l2.817-1.72Z"/>`),
		g.Group(children),
	)
}