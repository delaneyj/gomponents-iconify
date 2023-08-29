package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photomath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.087 33.828v7.586A2.082 2.082 0 0 1 35 43.5H12.143a2.082 2.082 0 0 1-2.088-2.086V6.586A2.082 2.082 0 0 1 12.143 4.5h22.856a2.082 2.082 0 0 1 2.088 2.086v7.774"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="4.72" d="m17.017 14.303l13.107 19.271m-13.107 0l13.107-19.271m4.514 5.734h9.307m-9.307 7.653h9.307"/>`),
		g.Group(children),
	)
}