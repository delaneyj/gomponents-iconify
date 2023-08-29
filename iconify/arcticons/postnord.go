package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postnord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM12.182 30.535v4.784m-.408-22.981h5.564m11.401.078h7.74m-24.577 4.761h9.945m-10.005 4.124h7.025m1.826-.035h1.317m-8.166 9.269v4.784m4.885-4.784v4.784m5.179-4.784v4.784m1.954-4.784v4.784m3.909-4.784v4.784m1.954-4.784v4.784m3.909-4.784v4.784m-13.289-4.784v4.784"/>`),
		g.Group(children),
	)
}