package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopSignFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.586 29h-9.172A1.986 1.986 0 0 1 10 28.414L3.586 22A1.986 1.986 0 0 1 3 20.586v-9.172A1.986 1.986 0 0 1 3.586 10L10 3.586A1.986 1.986 0 0 1 11.414 3h9.172A1.986 1.986 0 0 1 22 3.586L28.414 10A1.986 1.986 0 0 1 29 11.414v9.172A1.986 1.986 0 0 1 28.414 22L22 28.414a1.986 1.986 0 0 1-1.414.586Z"/>`),
		g.Group(children),
	)
}