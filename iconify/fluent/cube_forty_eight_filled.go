package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.679 5.332a6.25 6.25 0 0 1 4.642 0l15.007 6.003A4.25 4.25 0 0 1 44 15.281v17.438a4.25 4.25 0 0 1-2.672 3.946l-15.007 6.003a6.25 6.25 0 0 1-4.642 0L6.672 36.665A4.25 4.25 0 0 1 4 32.719V15.28a4.25 4.25 0 0 1 2.672-3.946l15.007-6.003Zm-7.426 10.273a1.25 1.25 0 0 0-1.005 2.29l9.502 4.171v9.184a1.25 1.25 0 1 0 2.5 0v-9.184l9.502-4.171a1.25 1.25 0 0 0-1.005-2.29L24 19.885l-9.748-4.28Z"/>`),
		g.Group(children),
	)
}