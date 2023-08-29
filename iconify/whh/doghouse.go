package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doghouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004.5 555.5Q985 575 957.5 575T911 556l-15-15v418q0 27-18.5 45.5T832 1023H704V767q0-79-56-135t-135.5-56T377 632t-56 135v256H192q-27 0-45.5-18.5T128 959V541l-15 15q-19 19-46.5 19t-47-19.5t-19.5-47T19 462L463 19q20-20 49-19q29-1 49 19l444 443q19 19 19 46.5t-19.5 47z"/>`),
		g.Group(children),
	)
}