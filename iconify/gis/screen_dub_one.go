package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenDubOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M23.836 35.625h4.453v28.75h-5.488V43.691c-2.005 1.875-4.369 3.262-7.09 4.16v-4.98c1.432-.469 2.988-1.354 4.668-2.656c1.68-1.315 2.832-2.845 3.457-4.59zM57 17.5h38v65H57Zm-9.5 0h5v65h-5zm-45-5A2.5 2.5 0 0 0 0 15v70a2.5 2.5 0 0 0 2.5 2.5h95A2.5 2.5 0 0 0 100 85V15a2.5 2.5 0 0 0-2.5-2.5z" color="currentColor"/>`),
		g.Group(children),
	)
}