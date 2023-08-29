package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gravitybox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.814 15.186c-2.044 2.814-1.017 4.23 2.293 3.155a12.47 12.47 0 0 1 1.205 3.709h0c-3.308 1.075-3.31 2.825 0 3.9h0a12.464 12.464 0 0 1-1.205 3.71c-3.309-1.076-4.338.339-2.293 3.154h0a12.467 12.467 0 0 1-3.155 2.293c-2.044-2.815-3.709-2.275-3.709 1.205h0a12.467 12.467 0 0 1-3.9 0h0c0-3.479-1.664-4.02-3.71-1.205a12.464 12.464 0 0 1-3.154-2.293c2.044-2.814 1.017-4.23-2.293-3.155a12.47 12.47 0 0 1-1.205-3.709h0c3.308-1.075 3.31-2.825 0-3.9a12.464 12.464 0 0 1 1.205-3.71c3.309 1.076 4.338-.339 2.293-3.154a12.467 12.467 0 0 1 3.155-2.293c2.045 2.814 3.709 2.275 3.709-1.205a12.467 12.467 0 0 1 3.9 0h0c0 3.479 1.664 4.02 3.71 1.205a12.464 12.464 0 0 1 3.154 2.293Z"/><circle cx="24" cy="24" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="16.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}