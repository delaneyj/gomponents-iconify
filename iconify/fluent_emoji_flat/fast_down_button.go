package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastDownButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="m15.038 17l-6.069-6.307C8.358 10.058 8.81 9 9.69 9h12.62c.882 0 1.332 1.058.72 1.693L16.963 17h5.348c.882 0 1.332 1.058.72 1.693l-6.31 6.558a1 1 0 0 1-1.44 0l-6.31-6.558C8.357 18.058 8.807 17 9.69 17h5.348Z"/></g>`),
		g.Group(children),
	)
}