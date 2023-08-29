package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastDownButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M8.97 10.693L15.037 17H9.69c-.882 0-1.332 1.058-.72 1.693l6.31 6.558a1 1 0 0 0 1.44 0l6.31-6.558c.612-.635.162-1.693-.72-1.693h-5.348l6.069-6.307c.611-.635.16-1.693-.72-1.693H9.69c-.882 0-1.332 1.058-.72 1.693Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}