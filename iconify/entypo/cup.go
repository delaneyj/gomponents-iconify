package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 1C5.721 1 3.06 2.41 3.205 3.555l1.442 13.467c.058.46 2.221 1.976 5.353 1.978c3.131-.002 5.295-1.518 5.351-1.979l1.442-13.467C16.938 2.41 14.279 1 10 1zm0 4.291c-3.132-.002-5.353-1.117-5.353-1.535C4.646 3.342 6.869 2.225 10 2.227c3.131-.002 5.354 1.115 5.351 1.529c0 .418-2.22 1.533-5.351 1.535z"/>`),
		g.Group(children),
	)
}