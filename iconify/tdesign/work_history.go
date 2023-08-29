package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkHistory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 2.5h9v4H22V11h-2V8.5H4v11h7v2H2v-15h5.5v-4Zm2 4h5v-2h-5v2Zm8.5 8a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm6.5-2.248v1.834L20.414 19L19 20.414l-2-2v-2.662h2Z"/>`),
		g.Group(children),
	)
}