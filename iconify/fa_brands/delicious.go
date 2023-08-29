package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delicious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M446.5 68c-.4-1.5-.9-3-1.4-4.5c-.9-2.5-2-4.8-3.3-7.1c-1.4-2.4-3-4.8-4.7-6.9c-2.1-2.5-4.4-4.8-6.9-6.8c-1.1-.9-2.2-1.7-3.3-2.5c-1.3-.9-2.6-1.7-4-2.4c-1.8-1-3.6-1.8-5.5-2.5c-1.7-.7-3.5-1.3-5.4-1.7c-3.8-1-7.9-1.5-12-1.5H48C21.5 32 0 53.5 0 80v352c0 4.1.5 8.2 1.5 12c2 7.7 5.8 14.6 11 20.3c1 1.1 2.1 2.2 3.3 3.3c5.7 5.2 12.6 9 20.3 11c3.8 1 7.9 1.5 12 1.5h352c26.5 0 48-21.5 48-48V80c-.1-4.1-.6-8.2-1.6-12zM416 432c0 8.8-7.2 16-16 16H224V256H32V80c0-8.8 7.2-16 16-16h176v192h192z"/>`),
		g.Group(children),
	)
}