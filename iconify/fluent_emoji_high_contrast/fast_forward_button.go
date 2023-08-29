package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m17 15.038l-6.307-6.069C10.058 8.358 9 8.81 9 9.69v12.62c0 .882 1.058 1.332 1.693.72L17 16.963v5.348c0 .882 1.058 1.332 1.693.72l6.558-6.31a1 1 0 0 0 0-1.44l-6.558-6.31C18.058 8.357 17 8.807 17 9.69v5.348Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}