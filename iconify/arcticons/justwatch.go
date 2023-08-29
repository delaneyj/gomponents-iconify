package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Justwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.483 40.033l-5.698 3.285a1.372 1.372 0 0 1-2.053-1.191v-6.57l7.751 4.477Zm0-10.695l-7.741 4.476v-8.953l7.741 4.477zm0-10.676l-7.741 4.477v-8.953l7.741 4.476zm0-10.696L8.795 4.683a1.372 1.372 0 0 0-2.053 1.191v6.57l7.741-4.477Zm19.809 21.372l-7.751 4.476v-8.953l7.751 4.477zm0-10.676l-7.751 4.477v-8.953l7.751 4.476zM24.318 24l-7.752 4.477v-8.954L24.318 24zm0-10.676l-7.752 4.457V8.848l7.752 4.476zm0 21.352l-7.752-4.457v8.933l7.752-4.476zM42.915 24.59l-6.72 3.877v-8.953l6.72 3.875c.47.28.47.942 0 1.202Z"/>`),
		g.Group(children),
	)
}