package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBallon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHeartBallon0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M17.333 4C13.333 4 12 7.156 12 11.05C12 18.1 19.8 24.51 24 26c4.2-1.49 12-7.9 12-14.95C36 7.156 34.312 4 30.667 4C28.434 4 25.194 7.077 24 8.889C22.806 7.077 19.566 4 17.333 4Z"/><path d="M24 26c-2 1.09-5 3.5-5 7s10 2.5 10 6s-11 5-11 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHeartBallon0)"/>`),
		g.Group(children),
	)
}