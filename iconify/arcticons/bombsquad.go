package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bombsquad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.996 12.409C29.498 10.804 37.69 16.4 39.294 24.91c1.604 8.51-3.989 16.71-12.49 18.314h-.001c-8.502 1.606-16.694-3.991-18.298-12.5c-1.604-8.51 3.988-16.71 12.49-18.315"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m21.038 12.297l1.625-4.933c4.781-.092 8.258.792 12.337 3.35l-1.118 4.843"/><path d="M29.249 8.07c2.333-6.49 8.05-2.954 10.52 1.328M19.046 19.461v2.461m8.299-3.314l.747 2.427"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.35 26.79c.664 3.359-1.123 6.728-4.557 8.592c-3.433 1.864-7.87 1.873-11.315.023c-3.445-1.85-5.252-5.21-4.609-8.573m19.804-.204h1.59m-20.615.25l-1.593-.428"/>`),
		g.Group(children),
	)
}