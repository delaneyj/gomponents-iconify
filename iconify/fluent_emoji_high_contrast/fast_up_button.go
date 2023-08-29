package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastUpButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16.962 15h5.348c.882 0 1.332-1.058.72-1.693l-6.31-6.558a1 1 0 0 0-1.44 0l-6.31 6.558c-.612.635-.162 1.693.72 1.693h5.348l-6.069 6.307c-.611.635-.16 1.693.72 1.693H22.31c.882 0 1.332-1.058.72-1.693L16.963 15Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}