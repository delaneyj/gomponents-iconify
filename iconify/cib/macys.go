package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Macys(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.021.833L12.255 12.39H0l9.927 7.177l-3.76 11.568L16 23.979l9.896 7.193l-3.781-11.62L32 12.391H19.786L16.02.828z"/>`),
		g.Group(children),
	)
}