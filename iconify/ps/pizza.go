package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pizza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m2 400l21 9q100 42 201 42t201-42l21-9L224 5zm248-262q-26 10-26 38q0 17 12.5 30t30.5 13q15 0 23-9l58 103q-61 30-126 30T98 313L224 91zM62 379l15-28q74 36 145 36q75 0 145-36l15 28q-75 30-158.5 30T62 379zm183-96q0 17-12.5 29.5T203 325q-18 0-30.5-12.5T160 283q0-18 12.5-30.5T203 240q17 0 29.5 12.5T245 283z"/>`),
		g.Group(children),
	)
}