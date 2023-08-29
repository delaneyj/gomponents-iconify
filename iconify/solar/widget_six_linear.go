package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidgetSixLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M2.5 6.5a4 4 0 1 1 8 0v2.667c0 .31 0 .465-.034.592a1 1 0 0 1-.707.707c-.127.034-.282.034-.592.034H6.5a4 4 0 0 1-4-4Zm11 8.333c0-.31 0-.465.034-.592a1 1 0 0 1 .707-.707c.127-.034.282-.034.592-.034H17.5a4 4 0 1 1-4 4v-2.667ZM2.5 17.5a4 4 0 0 1 4-4h2.4c.56 0 .84 0 1.054.109a1 1 0 0 1 .437.437c.109.214.109.494.109 1.054v2.4a4 4 0 0 1-8 0Zm11-11a4 4 0 1 1 4 4h-2.857c-.133 0-.2 0-.255-.006a1 1 0 0 1-.882-.882c-.006-.056-.006-.122-.006-.255V6.5Z"/>`),
		g.Group(children),
	)
}