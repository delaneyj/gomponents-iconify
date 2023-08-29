package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><radialGradient id="logosMoon0" cx="50%" cy="50%" r="49.789%" fx="50%" fy="50%"><stop offset="0%" stop-color="#4600D1"/><stop offset="49.285%" stop-color="#4600D1"/><stop offset="100%" stop-color="#35009F"/></radialGradient><radialGradient id="logosMoon1" cx="50%" cy="50%" r="49.603%" fx="50%" fy="50%"><stop offset="0%" stop-color="#35019E"/><stop offset="18.73%" stop-color="#320194"/><stop offset="100%" stop-color="#206"/></radialGradient><circle id="logosMoon2" cx="128" cy="128" r="128"/></defs><circle cx="128" cy="128" r="128" fill="#5805FF"/><mask id="logosMoon3" fill="#fff"><use href="#logosMoon2"/></mask><circle cx="199.694" cy="105.369" r="128" fill="url(#logosMoon0)" mask="url(#logosMoon3)"/><circle cx="275.372" cy="82.376" r="128" fill="url(#logosMoon1)" mask="url(#logosMoon3)"/>`),
		g.Group(children),
	)
}