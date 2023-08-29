package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M20.888 15.481a5.25 5.25 0 1 0-5.8 2.863l-1.892 3.085a1.75 1.75 0 0 0 2.983 1.83l4.532-7.391c.076-.124.135-.254.177-.387Zm-2.982-2.278a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Z"/></g>`),
		g.Group(children),
	)
}