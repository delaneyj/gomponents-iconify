package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.942 12.57L10 4.335V1h.5c.275 0 .5-.225.5-.5s-.225-.5-.5-.5h-5c-.275 0-.5.225-.5.5s.225.5.5.5H6v3.335L1.058 12.57C-.074 14.456.8 16 3 16h10c2.2 0 3.074-1.543 1.942-3.43zM3.766 10L7 4.61V1h2v3.61L12.234 10H3.766z"/>`),
		g.Group(children),
	)
}