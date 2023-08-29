package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackberrymessenger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.47 7.28H6.53a2 2 0 0 0-2 1.88v22.25a2 2 0 0 0 2 1.88h2.61v7.43l8.25-7.43h24.08a2 2 0 0 0 2-1.88V9.16a2 2 0 0 0-2-1.88ZM12.05 17.91h3.4c2.2 0 2.73 1.09 2.73 1.95c0 1.21-.79 2.53-3.56 2.53h-3.54Zm3.8-2.25h-3.54l1-4.47h3.4c2.21 0 2.74 1.08 2.74 1.94c-.04 1.22-.83 2.53-3.6 2.53Zm6.61 13.73h-3.54l1-4.48h3.4C25.5 24.91 26 26 26 26.86c0 1.22-.77 2.53-3.54 2.53Zm1.32-7h-3.54l1-4.48h3.41c2.2 0 2.73 1.09 2.73 1.95c-.04 1.21-.83 2.53-3.6 2.53ZM25 15.66h-3.53l1-4.47h3.4c2.21 0 2.73 1.08 2.73 1.94c-.04 1.22-.82 2.53-3.6 2.53Zm7 11h-3.5l1-4.47h3.4c2.21 0 2.73 1.09 2.73 1.94c-.04 1.22-.82 2.53-3.63 2.53Zm1.32-7h-3.5l1-4.48h3.4c2.2 0 2.73 1.09 2.73 1.95c-.03 1.22-.82 2.53-3.59 2.53Z"/>`),
		g.Group(children),
	)
}