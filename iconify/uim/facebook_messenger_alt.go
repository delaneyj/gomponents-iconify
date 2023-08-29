package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookMessengerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a9.651 9.651 0 0 0-10 9.7a9.49 9.49 0 0 0 3.14 7.175a.806.806 0 0 1 .27.57l.055 1.779a.801.801 0 0 0 1.122.708l1.984-.875a.798.798 0 0 1 .534-.04A10.876 10.876 0 0 0 12 21.4A9.705 9.705 0 1 0 12 2Z" opacity=".5"/><path fill="currentColor" d="M6.499 14.772a1 1 0 0 1-.765-1.642l3.052-3.636a.996.996 0 0 1 1.29-.21l3.346 2.056l2.312-2.755a1 1 0 1 1 1.532 1.285l-2.867 3.416a1 1 0 0 1-1.289.21L9.764 11.44l-2.498 2.975a.994.994 0 0 1-.767.357Z"/>`),
		g.Group(children),
	)
}