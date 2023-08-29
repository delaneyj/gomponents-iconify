package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoNpm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 28V4h24v24H4zM8.5 8.5v15H16v-12h4.5v12h3v-15h-15z"/>`),
		g.Group(children),
	)
}