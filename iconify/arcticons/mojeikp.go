package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mojeikp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13.458c1.376-3.022 3.636-5.03 6.88-5.801c3.263-.777 6.264-.041 8.746 2.164c3.596 3.196 4.17 7.87 2.42 11.649c-1.353 2.916-3.436 5.298-5.591 7.603c-2.579 2.759-5.344 5.336-8.174 7.839A583.959 583.959 0 0 1 24 40.643a592.324 592.324 0 0 1-4.28-3.731c-2.83-2.503-5.597-5.08-8.175-7.839c-2.155-2.305-4.239-4.687-5.59-7.603c-1.752-3.778-1.177-8.453 2.419-11.648c2.482-2.206 5.483-2.942 8.746-2.165c3.244.772 5.504 2.779 6.88 5.801Zm0 0v27.185"/>`),
		g.Group(children),
	)
}