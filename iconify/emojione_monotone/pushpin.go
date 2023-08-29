package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M62 23.448L40.544 2c-.315 2.875.297 6.054 1.764 9.111L27.77 25.646c-5.529-3.519-11.593-5.079-16.941-4.483l14.143 14.148C18.714 42.195 4.687 59.055 2.59 61.152l-.071.068c-.845.844-.586 1.106.258.262c1.852-1.846 18.983-16.113 25.92-22.445L42.83 53.174c.592-5.298-.927-11.291-4.369-16.772l14.62-14.616c2.995 1.399 6.103 1.975 8.919 1.662"/>`),
		g.Group(children),
	)
}