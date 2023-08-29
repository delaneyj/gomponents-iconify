package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#B68900"/><path fill="#FFF" d="M26.633 15.817L15.817 5L5 15.817l10.817 3.996l10.816-3.996zM8.364 14.9l7.333-7.498s7.169 7.333 7.471 7.48c.303.146-4.931 0-4.931 0l-2.42-2.475l-2.448 2.493H8.364zm7.453 5.674l10.816-4.024l-10.816 10.083L5 16.605l10.817 3.97z"/></g>`),
		g.Group(children),
	)
}