package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Detexify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 8a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2Zm15.6 23.98h5.77M21.6 20.44h5.77m-5.77 5.77h3.76m-3.76-5.77v11.54"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.07 17.55l3.88 5.79l-3.88 5.78m7.75-11.57l-3.87 5.79l3.87 5.78M10.18 17.56h8.71m-4.35 11.55V17.56m4.35 0v1.56m-8.71-1.56v1.56"/>`),
		g.Group(children),
	)
}