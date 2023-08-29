package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.5a2 2 0 0 1 2-2H25A2.5 2.5 0 0 1 27.5 4v23.5H7a.5.5 0 0 0 0 1h20.622l-.236.667A2 2 0 0 1 25.5 30.5h-19a2 2 0 0 1-2-2v-25Zm20.291 22l1.709-1.709v-2.878L21.913 25.5h2.878ZM8.5 3.912v2.88L12.791 2.5H9.912L8.5 3.912ZM14.912 2.5L8.5 8.912v2.88L17.791 2.5h-2.879Zm5 0L8.5 13.912v2.88L22.791 2.5h-2.878Zm5 0L8.5 18.913v2.878L26.487 3.804A1.5 1.5 0 0 0 25 2.5h-.087ZM26.5 5.912l-18 18V25.5h1.291L26.5 8.791V5.912Zm0 5L11.912 25.5h2.88L26.5 13.791v-2.879Zm0 5L16.913 25.5h2.878l6.709-6.709v-2.879Z"/>`),
		g.Group(children),
	)
}