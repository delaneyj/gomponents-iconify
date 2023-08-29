package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashingAway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.976 10C22.7 7.197 20.13 5 17 5s-5.7 2.197-5.976 5H11c-.464 2.553-2.436 4.028-4.416 4.877a24.653 24.653 0 0 0 4.779-.858a.5.5 0 0 1 .274.962C8.14 15.98 5.132 15.999 2.66 16A1.2 1.2 0 0 0 2 17.056c0 .39.2.739.51.944h.003c2.508 0 5.62 0 9.235 2.066a.5.5 0 0 1-.496.868a13.933 13.933 0 0 0-4.565-1.666C8.662 20.126 10.453 21.54 11 24c.014.06.028.118.044.173C11.406 26.893 13.934 29 17 29c3.13 0 5.7-2.197 5.976-5H23a7 7 0 1 0 0-14h-.024ZM23 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6 1a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-5.5 22a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}