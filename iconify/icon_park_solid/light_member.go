package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightMember(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLightMember0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="M35.056 15L18.463 7.665a2 2 0 0 0-2.447.682L11.359 15"/><path fill="#fff" stroke="#fff" d="M43 15H5a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h38a1 1 0 0 0 1-1V16a1 1 0 0 0-1-1Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m19 23l5.103 10L29 23"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLightMember0)"/>`),
		g.Group(children),
	)
}