package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.845 9.983L9.88 27.336c0 .977 2.74 1.77 6.12 1.77s6.12-.793 6.12-1.77L24.5 9.85c-2.455 1.024-6.812 1.134-8.498 1.134c-1.61 0-5.655-.1-8.155-1zm16.285-4.23l-.376-1.68c0-.65-3.472-1.178-7.754-1.178s-7.754.53-7.754 1.18L7.87 5.752c-.714.284-1.12.608-1.12.953V7.99c0 1.1 4.142 1.994 9.25 1.994s9.25-.894 9.25-1.995V6.704c0-.345-.406-.67-1.12-.953z"/>`),
		g.Group(children),
	)
}