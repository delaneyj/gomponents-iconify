package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func House(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004.5 555.5Q985 575 957.5 575T911 555l-15-14v418h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5H33q-13 0-22.5-9.5T1 991t9.5-22.5T33 959h95V541l-15 14q-19 20-46.5 20t-47-19.5t-19.5-47T19 462L463 19q20-20 49-19q29-1 49 19l443 443q20 19 20 46.5t-19.5 47zM448 671q0-13-9.5-22.5T416 639H288q-13 0-22.5 9.5T256 671v288h192V671zm320 0q0-13-9.5-22.5T736 639H608q-13 0-22.5 9.5T576 671v128q0 13 9.5 22.5T608 831h128q13 0 22.5-9.5T768 799V671z"/>`),
		g.Group(children),
	)
}