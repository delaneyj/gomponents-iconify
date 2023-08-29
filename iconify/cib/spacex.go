package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spacex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31.937 9.287c-.011 0-.016-.005-.02-.005a.068.068 0 0 0-.032.011C11.713 11.324 2.234 20.094 0 22.252l.297.468h3.525c9.161-9.213 21.541-12.271 28.089-13.276l.005.005c.004 0 .009-.011.015-.011a.078.078 0 0 0 .068-.077c0-.037-.027-.063-.063-.073zM.505 14.011l-.213.401l4.328 3.156a42.247 42.247 0 0 1 2.683-1.432l-2.901-2.125zm9.62 4.186a45.27 45.27 0 0 0-2.235 1.756l3.803 2.765h3.943l.167-.359z"/>`),
		g.Group(children),
	)
}