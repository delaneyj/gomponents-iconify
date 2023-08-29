package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataHistogramTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 5.23a2.25 2.25 0 0 1 2.25-2.25h2.5a2.25 2.25 0 0 1 2.25 2.25V7h3.25A2.25 2.25 0 0 1 21 9.25v11a.75.75 0 0 1-.75.75H3.75a.75.75 0 0 1-.75-.75v-8A2.25 2.25 0 0 1 5.25 10H8.5V5.23ZM10 19.5h4V5.23a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0-.75.75V19.5Zm-1.5-8H5.25a.75.75 0 0 0-.75.75v7.25h4v-8Zm7 8h4V9.25a.75.75 0 0 0-.75-.75H15.5v11Z"/>`),
		g.Group(children),
	)
}