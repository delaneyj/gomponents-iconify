package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartDecoration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 2a4 4 0 0 0-4 4v20a4 4 0 0 0 4 4h20a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4H6Zm10 9.54s1.41-2.17 3.24-2.49c4.54-.79 6.41 3.18 5.65 6.13c-1.35 5.27-8.89 9.87-8.89 9.87s-7.54-4.6-8.89-9.86c-.75-2.95 1.12-6.93 5.65-6.13c1.83.32 3.24 2.48 3.24 2.48Z"/>`),
		g.Group(children),
	)
}