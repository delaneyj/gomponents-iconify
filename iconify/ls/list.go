package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M129 118c0-36-28-64-64-64S0 82 0 118s29 65 65 65s64-29 64-65zm75 53h564V67H204v104zm-75 176c0-36-28-65-64-65S0 311 0 347s29 64 65 64s64-28 64-64zm75 52h564V295H204v104zm-75 176c0-36-28-64-64-64S0 539 0 575s29 65 65 65s64-29 64-65zm75 51h564V523H204v103z"/>`),
		g.Group(children),
	)
}