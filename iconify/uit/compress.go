package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 16h-5a.5.5 0 0 0 0 1H7v4.5a.5.5 0 1 0 1 0v-5a.5.5 0 0 0-.5-.5zm9-8h5a.5.5 0 0 0 0-1H17V2.5a.5.5 0 0 0-1 0v5a.5.5 0 0 0 .5.5zm-9-6a.5.5 0 0 0-.5.5V7H2.5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5zm14 14h-5a.5.5 0 0 0-.5.5v5a.5.5 0 1 0 1 0V17h4.5a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}