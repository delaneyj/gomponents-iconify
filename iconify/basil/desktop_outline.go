package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12 3.75l7.274.205a2.584 2.584 0 0 1 2.494 2.29a37.425 37.425 0 0 1 0 8.51a2.584 2.584 0 0 1-2.494 2.29l-5.524.156v2.05H17a.75.75 0 1 1 0 1.5H7a.75.75 0 0 1 0-1.5h3.25V17.2l-5.524-.156a2.584 2.584 0 0 1-2.494-2.29a37.434 37.434 0 0 1 0-8.51a2.584 2.584 0 0 1 2.494-2.29L12 3.75Zm0 1.5l-7.231.205c-.54.015-.986.424-1.047.96a35.934 35.934 0 0 0 0 8.17c.062.536.507.945 1.047.96l7.23.205l7.232-.205c.54-.015.985-.424 1.047-.96a35.94 35.94 0 0 0 0-8.17a1.084 1.084 0 0 0-1.047-.96L12 5.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}