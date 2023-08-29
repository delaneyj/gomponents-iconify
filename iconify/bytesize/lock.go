package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 15v15h22V15Zm4 0C9 9 9 5 16 5s7 4 7 10m-7 5v3"/><circle cx="16" cy="24" r="1"/></g>`),
		g.Group(children),
	)
}