package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 10.34a2.11 2.11 0 0 0-1.16-1.9a2 2 0 0 0-2.13.15L26 11.6V8a2 2 0 0 0-2-2H6a4 4 0 0 0-4 4v16a4 4 0 0 0 4 4h18a2 2 0 0 0 2-2v-3.6l4.64 3a2.07 2.07 0 0 0 2.2.2A2.11 2.11 0 0 0 34 25.66Zm-2.07 15.43c-.06 0-.11 0-.19-.06L24 20.77V28H6a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h18v7.23l7.8-5a.11.11 0 0 1 .13 0a.11.11 0 0 1 .07.11v15.32a.11.11 0 0 1-.07.11Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}