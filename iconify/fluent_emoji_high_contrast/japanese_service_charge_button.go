package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseServiceChargeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M21 7a1 1 0 0 1 1 1v2.875c0 .069.056.125.125.125H25a1 1 0 1 1 0 2h-2.866a.126.126 0 0 0-.125.133c.223 4.32-.678 7.35-2.036 9.522c-1.367 2.188-3.14 3.41-4.444 4.194a1 1 0 1 1-1.028-1.716c1.192-.716 2.653-1.741 3.776-3.538c1.122-1.796 1.965-4.458 1.725-8.536l-.002-.03a.03.03 0 0 0-.03-.029h-7.845a.125.125 0 0 0-.125.125V18a1 1 0 1 1-2 0v-4.875A.125.125 0 0 0 9.875 13H7a1 1 0 1 1 0-2h2.875a.125.125 0 0 0 .125-.125V8a1 1 0 1 1 2 0v2.875c0 .069.056.125.125.125h7.75a.125.125 0 0 0 .125-.125V8a1 1 0 0 1 1-1Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}