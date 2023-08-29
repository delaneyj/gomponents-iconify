package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHarm0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M6 9.256L24.009 4L42 9.256v10.778A26.316 26.316 0 0 1 24.003 45A26.32 26.32 0 0 1 6 20.029V9.256Z"/><path stroke="#000" stroke-linecap="round" d="M29.5 18.408L18.186 29.722m0-11.314L29.5 29.722"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHarm0)"/>`),
		g.Group(children),
	)
}