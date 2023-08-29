package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M201.524 512L96.134 264.33H206.63l-56.373-129.967l-87.19 181.587H0L150.257 0l135.334 315.95H172.336L201.524 512z"/>`),
		g.Group(children),
	)
}