package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.694 11.343c-.615 0-1.15.356-1.403.882a1.75 1.75 0 0 1-3.153-1.52a5.056 5.056 0 1 1 8.475 5.388a5.056 5.056 0 1 1-8.737 4.728a1.75 1.75 0 1 1 3.335-1.061a1.557 1.557 0 1 0 1.93-1.964a1.818 1.818 0 0 1-.087-.023a1.75 1.75 0 0 1 .252-3.443a1.558 1.558 0 0 0-.612-2.987Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}