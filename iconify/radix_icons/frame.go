package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 1.5a.5.5 0 0 0-1 0V4H5V1.5a.5.5 0 0 0-1 0V4H1.5a.5.5 0 0 0 0 1H4v5H1.5a.5.5 0 0 0 0 1H4v2.5a.5.5 0 0 0 1 0V11h5v2.5a.5.5 0 0 0 1 0V11h2.5a.5.5 0 0 0 0-1H11V5h2.5a.5.5 0 0 0 0-1H11V1.5ZM10 10V5H5v5h5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}