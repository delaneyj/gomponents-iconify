package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M10.5 3c-.709 0-1.18.288-1.406.688L7.813 6H6c0-.551-.449-1-1-1H3c-.551 0-1 .449-1 1v.188A2.985 2.985 0 0 0 0 9v10.125A4.874 4.874 0 0 0 4.875 24h16.25A4.874 4.874 0 0 0 26 19.125v-8.25A4.874 4.874 0 0 0 21.125 6h-.938l-1.28-2.313C18.68 3.29 18.18 3 17.5 3h-7zM4 6.5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 4 6.5zm10 .438a7.063 7.063 0 1 1-.001 14.126A7.063 7.063 0 0 1 14 6.937zm0 2.25a4.812 4.812 0 1 0 0 9.624a4.812 4.812 0 0 0 0-9.625z"/>`),
		g.Group(children),
	)
}