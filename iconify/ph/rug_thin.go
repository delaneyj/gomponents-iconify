package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RugThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 20a4 4 0 0 0-4 4v20h-40V24a4 4 0 0 0-8 0v20h-40V24a4 4 0 0 0-8 0v20H60V24a4 4 0 0 0-8 0v208a4 4 0 0 0 8 0v-20h40v20a4 4 0 0 0 8 0v-20h40v20a4 4 0 0 0 8 0v-20h40v20a4 4 0 0 0 8 0V24a4 4 0 0 0-4-4ZM60 52h136v152H60Zm68 120a4 4 0 0 0 3.43-1.94l24-40a4 4 0 0 0 0-4.12l-24-40a4 4 0 0 0-6.86 0l-24 40a4 4 0 0 0 0 4.12l24 40A4 4 0 0 0 128 172Zm0-76.23L147.33 128L128 160.23L108.67 128Z"/>`),
		g.Group(children),
	)
}