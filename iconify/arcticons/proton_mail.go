package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 9.266a.896.896 0 0 1 1.47-.69L21.421 21.41a4.035 4.035 0 0 0 5.156 0L42.03 8.576a.896.896 0 0 1 1.469.69v26.332a4.035 4.035 0 0 1-4.035 4.034H8.535A4.035 4.035 0 0 1 4.5 35.598Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.265 19.178l.002.002l-7.782 6.874a3.586 3.586 0 0 1-4.667.07L4.5 15.879m30.707-1.636v25.389"/>`),
		g.Group(children),
	)
}