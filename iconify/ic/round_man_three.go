package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundManThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 7h-4c-1.1 0-2 .9-2 2v5c0 .55.45 1 1 1h1v6c0 .55.45 1 1 1h2c.55 0 1-.45 1-1v-6h1c.55 0 1-.45 1-1V9c0-1.1-.9-2-2-2zm-2.35-1.1L10.1 4.35c-.2-.2-.2-.51 0-.71l1.54-1.54c.2-.2.51-.2.71 0l1.54 1.54c.2.2.2.51 0 .71L12.35 5.9c-.19.19-.51.19-.7 0z"/>`),
		g.Group(children),
	)
}