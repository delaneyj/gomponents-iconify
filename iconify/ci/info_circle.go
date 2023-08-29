package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10ZM9.99 10.99V13h1v4h3.02v-2H13l.01-4.009l-3.02-.001Zm1-3.99v2.019h2.02V7h-2.02Z"/>`),
		g.Group(children),
	)
}