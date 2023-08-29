package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 193q9.5 0 16.5 6.5t7 16.5t-7 16.5t-16.5 6.5t-16.5-6.5t-7-16.5t7-16.5t16.5-6.5zm0-190q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM260 263l81-175l-174 81l-82 175z"/>`),
		g.Group(children),
	)
}