package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Information(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><g fill="#fff"><path d="M27 27.8h10v24H27z"/><circle cx="32" cy="17.2" r="5"/></g>`),
		g.Group(children),
	)
}