package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 12.5L11 15l4.5-4.5m-.595-5.512l-.48-.659a3 3 0 0 0-4.85 0l-.48.659l-.804-.127a3 3 0 0 0-3.43 3.43l.127.804l-.659.48a3 3 0 0 0 0 4.85l.659.48l-.127.804a3 3 0 0 0 3.43 3.43l.804-.127l.48.659a3 3 0 0 0 4.85 0l.48-.659l.804.127a3 3 0 0 0 3.43-3.43l-.127-.804l.659-.48a3 3 0 0 0 0-4.85l-.659-.48l.127-.804a3 3 0 0 0-3.43-3.43l-.804.127z"/>`),
		g.Group(children),
	)
}