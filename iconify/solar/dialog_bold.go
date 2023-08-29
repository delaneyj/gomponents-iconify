package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialogBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 8.5a6.5 6.5 0 0 0-11.627-3.993a9.5 9.5 0 0 1 9.121 9.823l.333.088a1.3 1.3 0 0 0 1.591-1.591l-.127-.476c-.103-.385-.04-.791.125-1.153A6.477 6.477 0 0 0 22 8.5Z"/><path fill-rule="evenodd" d="M18 14a8 8 0 0 1-11.45 7.22a1.671 1.671 0 0 0-1.15-.13l-1.227.329a1.3 1.3 0 0 1-1.591-1.592L2.91 18.6a1.67 1.67 0 0 0-.13-1.15A8 8 0 1 1 18 14ZM6.5 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}