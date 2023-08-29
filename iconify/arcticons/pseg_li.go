package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsegLi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.721 24l6.78-3.355l-16.065 1.899l14.861-6.388l-7.548.484l4.193-6.296l-12.963 9.677l9.677-12.963l-6.296 4.193l.484-7.548l-6.388 14.862l1.9-16.065L24 9.28L20.645 2.5l1.898 16.065l-6.387-14.862l.483 7.548l-6.295-4.193l9.677 12.963l-12.963-9.677l4.193 6.296l-7.548-.484l14.861 6.388l-16.064-1.9L9.279 24l-6.78 3.355l16.065-1.899l-14.861 6.388l7.548-.484l-4.193 6.296l12.963-9.677l-9.677 12.963l6.296-4.193l-.484 7.548l6.387-14.862L20.645 45.5L24 38.72l3.355 6.78l-1.898-16.065l6.387 14.862l-.484-7.548l6.296 4.193l-9.677-12.963l12.963 9.677l-4.193-6.296l7.548.484l-14.861-6.388l16.064 1.9L38.721 24Z"/>`),
		g.Group(children),
	)
}