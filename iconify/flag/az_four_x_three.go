package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#3f9c35" d="M.1 0h640v480H.1z"/><path fill="#ed2939" d="M.1 0h640v320H.1z"/><path fill="#00b9e4" d="M.1 0h640v160H.1z"/><circle cx="304" cy="240" r="72" fill="#fff"/><circle cx="320" cy="240" r="60" fill="#ed2939"/><path fill="#fff" d="m384 200l7.7 21.5l20.6-9.8l-9.8 20.7L424 240l-21.5 7.7l9.8 20.6l-20.6-9.8L384 280l-7.7-21.5l-20.6 9.8l9.8-20.6L344 240l21.5-7.7l-9.8-20.6l20.6 9.8L384 200z"/>`),
		g.Group(children),
	)
}