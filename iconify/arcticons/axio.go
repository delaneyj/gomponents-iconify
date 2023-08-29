package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.293 23.962l-17.15 17.15c-1.716 1.716-4.54 1.716-6.256 0h0c-1.715-1.714-1.715-4.54 0-6.254l10.896-10.896L6.786 13.066c-1.715-1.715-1.715-4.54 0-6.254h0c1.715-1.716 4.54-1.716 6.255 0l17.252 17.15Z"/><circle cx="36.85" cy="11.15" r="5.65" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.872 34.152l6.254-6.255l7.062 7.062c1.715 1.715 1.715 4.54 0 6.255h0c-1.715 1.715-4.54 1.715-6.255 0l-7.061-7.062Z"/>`),
		g.Group(children),
	)
}