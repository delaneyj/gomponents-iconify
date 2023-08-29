package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBowl0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M32 44s0-3.864.002-3.872a17.964 17.964 0 0 0 4.726-3.4A17.943 17.943 0 0 0 42 24H6c0 4.97 2.015 9.47 5.272 12.728a18.085 18.085 0 0 0 4.741 3.407L16 44h16Z"/><path d="M24 18.008V8m12 10.008V12m-24 6.008V12m28-4a4 4 0 0 0-4 4m-8-8a4 4 0 0 0-4 4m-8 0a4 4 0 0 0-4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBowl0)"/>`),
		g.Group(children),
	)
}