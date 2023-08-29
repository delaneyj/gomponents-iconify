package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 240.2V256H0c0-20 10-38.7 26.6-49.8L274.9 40.7c8.6-5.7 18.6-8.7 28.9-8.7c115 0 208.2 93.2 208.2 208.2zm0 47.8v128c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V288h512z"/>`),
		g.Group(children),
	)
}