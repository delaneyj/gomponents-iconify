package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingGiftRewardBoxSocialPresentGiftMediaRatingBow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="5" x=".5" y="3" rx="1"/><path d="M12.5 8v4.5a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1V8M7 3v10.5m3-13L7 3L4 .5"/></g>`),
		g.Group(children),
	)
}