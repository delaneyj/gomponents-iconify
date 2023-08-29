package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialogBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M18 14a8 8 0 0 1-11.45 7.22a1.671 1.671 0 0 0-1.15-.13l-1.227.329a1.3 1.3 0 0 1-1.591-1.592L2.91 18.6a1.67 1.67 0 0 0-.13-1.15A8 8 0 1 1 18 14ZM6.5 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/><path d="M17.984 14.508a6.43 6.43 0 0 0 .32-.142c.291-.14.622-.189.934-.105l.996.267a1.056 1.056 0 0 0 1.294-1.294l-.267-.996a1.358 1.358 0 0 1 .105-.935A6.5 6.5 0 1 0 9.492 6.016a8 8 0 0 1 8.493 8.492Z" opacity=".6"/></g>`),
		g.Group(children),
	)
}