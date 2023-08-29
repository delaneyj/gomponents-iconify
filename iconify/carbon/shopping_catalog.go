package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCatalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 24h4v4h-4zm7 0h4v4h-4zm-7-7h4v4h-4zm7 0h4v4h-4z"/><path fill="currentColor" d="M28 11h-6V7c0-1.7-1.3-3-3-3h-6c-1.7 0-3 1.3-3 3v4H4c-.6 0-1 .4-1 1v.2l1.9 12.1c.1 1 1 1.7 2 1.7H15v-2H6.9L5.2 13H28v-2zM12 7c0-.6.4-1 1-1h6c.6 0 1 .4 1 1v4h-8V7z"/>`),
		g.Group(children),
	)
}