package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuebecFiveHundredEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.042 18.266h2.826m-2.826-5.653h2.826m-2.826 2.827h1.837m-1.837-2.827v5.653m-16.465 0l-1.837-1.837m-1.908-.07a1.873 1.873 0 1 0 3.745 0V14.45a1.9 1.9 0 0 0-1.908-1.908a1.841 1.841 0 0 0-1.837 1.908Zm5.3-3.746v3.745a1.839 1.839 0 1 0 3.675 0v-3.745m3.108-1.484l.99-.353m-2.473 7.49h2.826m-2.826-5.653h2.826m-2.826 2.827h1.837m-1.837-2.827v5.653m17.736-1.837a1.908 1.908 0 0 1-3.816 0V14.52a1.9 1.9 0 0 1 1.908-1.908a1.841 1.841 0 0 1 1.837 1.908M9.479 20.81H38.52M9.832 36.015c1.037.864 1.9 1.21 4.146 1.21h.518a4.491 4.491 0 1 0 0-8.983H9.832v-4.836h9.155m2.54 13.646h6.91m-6.91-11.919l3.455-1.9v13.82m6.277 0h6.909m-6.91-11.92l3.455-1.9v13.82M27.137 15.44a1.431 1.431 0 0 1 0 2.862h-2.36v-5.724h2.36a1.431 1.431 0 0 1 0 2.862Zm0 0h-2.361"/><rect width="35" height="39" x="6.5" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}