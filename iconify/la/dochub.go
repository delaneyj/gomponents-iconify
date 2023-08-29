package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dochub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 3.586V12h8.414L19 3.586zM7 4v24h11c4.411 0 8-3.589 8-8v-6h-6v4.938A3.066 3.066 0 0 1 16.937 22H13V10h4V4H7zm2 2h6v2h-4v16h5.938A5.068 5.068 0 0 0 22 18.937V16h2v4c0 3.309-2.691 6-6 6H9V6zm12 2.414L22.586 10H21V8.414z"/>`),
		g.Group(children),
	)
}