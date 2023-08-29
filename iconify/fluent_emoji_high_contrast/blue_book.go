package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.5a2 2 0 0 1 2-2H25A2.5 2.5 0 0 1 27.5 4v23.5H7a.5.5 0 0 0 0 1h20.622l-.236.667A2 2 0 0 1 25.5 30.5h-19a2 2 0 0 1-2-2v-25Zm20.587 22L8.5 8.912v2.88L22.209 25.5h2.878ZM8.5 3.912v2.88l18 18v-2.88l-18-18Zm18 15.88v-2.88L12.088 2.5h-2.88L26.5 19.791ZM14.209 2.5L26.5 14.791v-2.879L17.087 2.5H14.21Zm5 0L26.5 9.791V6.912L22.087 2.5H19.21Zm5 0L26.5 4.791V4A1.5 1.5 0 0 0 25 2.5h-.791Zm-4.122 23L8.5 13.912v2.88l8.709 8.708h2.878Zm-5 0L8.5 18.913v2.878l3.709 3.709h2.879Zm-5 0L8.5 23.913V25.5h1.588Z"/>`),
		g.Group(children),
	)
}