package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 26h24v2H4zm12-3a7 7 0 0 1-7-7V5h2v11a5 5 0 0 0 10 0V5h2v11a7 7 0 0 1-7 7z"/>`),
		g.Group(children),
	)
}