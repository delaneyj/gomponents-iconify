package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagVn1x10"><path fill-opacity=".7" d="M177.2 0h708.6v708.7H177.2z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagVn1x10)" transform="translate(-128) scale(.72249)"><path fill="#da251d" d="M0 0h1063v708.7H0z"/><path fill="#ff0" d="m661 527.5l-124-92.6l-123.3 93.5l45.9-152l-123.2-93.8l152.4-1.3L536 129.8L584.3 281l152.4.2l-122.5 94.7L661 527.5z"/></g>`),
		g.Group(children),
	)
}