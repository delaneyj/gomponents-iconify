package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ppt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 11h3v12h2V11h3V9h-8v2zm-8 12h-2V9h6a2.002 2.002 0 0 1 2 2v5a2.002 2.002 0 0 1-2 2h-4zm0-7h4v-5.002h-4zM4 23H2V9h6a2.002 2.002 0 0 1 2 2v5a2.002 2.002 0 0 1-2 2H4zm0-7h4v-5.002H4z"/>`),
		g.Group(children),
	)
}