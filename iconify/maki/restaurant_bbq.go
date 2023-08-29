package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestaurantBbq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.017 12.5H14.5L12 15v-2.483a1.619 1.619 0 0 1-1.137.48a1.6 1.6 0 0 1-.993-.345a6.784 6.784 0 0 1 .008-10.31a1.618 1.618 0 0 1 2.122.12V0l2.5 2.5h-2.464a1.593 1.593 0 0 1 .122.129a1.624 1.624 0 0 1-.287 2.28a3.515 3.515 0 0 0-1.121 2.586a3.558 3.558 0 0 0 1.147 2.616a1.635 1.635 0 0 1 .232 2.279c-.034.041-.075.072-.112.11ZM2.504 0l-1 5.5c-.146.805 1.781 1.181 1.75 2l-.25 6.5a.963.963 0 0 0 1 1a.963.963 0 0 0 1-1l-.25-6.5c-.031-.818 1.733-1.18 1.75-2l-1-5.5h-.5l.25 4l-.75.5l-.25-4.5h-.5l-.25 4.5l-.75-.5l.25-4Z"/>`),
		g.Group(children),
	)
}