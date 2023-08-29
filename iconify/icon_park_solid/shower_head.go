package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowerHead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShowerHead0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M27 20v2H9v-2c0-3.314 4.03-6 9-6s9 2.686 9 6Z"/><path d="M42 44V12.5C42 7.806 36.627 4 30 4s-12 3.806-12 8.5V14m0 15v-1m-7.829.03l-.342.94M4.171 43.03l-.342.94M18 44v-1m0-6v-2m-10.658.06l-.684 1.88"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShowerHead0)"/>`),
		g.Group(children),
	)
}