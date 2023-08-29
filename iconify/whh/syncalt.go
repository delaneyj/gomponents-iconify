package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syncalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1012 320H897v384q0 26-18.5 44.5T833 767H705q-27 0-45.5-18.5T641 704V320H527q-9 0-12.5-18t4.5-29L747 11q9-11 22-11t22 11l228 262q8 11 5 29t-12 18zm-733 693q-9 11-22 11t-22-11L7 751q-9-11-5.5-29T14 704h115V320q0-27 18.5-45.5T193 256h128q27 0 45.5 18.5T385 320v384h114q9 0 12.5 18t-4.5 29z"/>`),
		g.Group(children),
	)
}