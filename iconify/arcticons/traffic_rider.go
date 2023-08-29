package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficRider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.618" cy="29.507" r="4.882" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.618" cy="29.507" r="1.797" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.382" cy="29.507" r="4.882" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.382" cy="29.507" r="1.797" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.809 28.422l5.542-6.219m21.641 5.623l-5.975-14.215h-3.398M7.586 21.383l12.841 6.443m3.251-1.756L8.582 18.102l3.594-.937l8.994 2.152"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.909 19.088s1.21.577 1.913 2.06h5.39l-1.123-4.032h-6.932c-1.738 0-3.184 1.621-3.32 3.602h5.753m.34 6.924l5.282-6.493m-1.318 4.491c.098-1.65 3.535-7.167 9.686-4.667"/><ellipse cx="23.678" cy="27.826" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.252" ry="1.756"/>`),
		g.Group(children),
	)
}