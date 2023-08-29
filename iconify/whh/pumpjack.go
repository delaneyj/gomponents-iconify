package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pumpjack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.054 1024h-960q-13 0-22.5-9.5T.054 992v-64q0-13 9.5-22.5t22.5-9.5h160l210-590l-183-18l-27 192q0 13-9.5 22.5t-22.5 9.5q-26 0-47.5-5t-37-12t-28-21t-20-26t-13-33t-8.5-35t-4.5-40.5t-1.5-41V256q0-61 36-122t88-97.5t100-36.5q13 0 22.5 9.5t9.5 22.5l-24 167l167 17q19-24 49-24q38 0 56 34l364 37q12 1 21 25.5t7 38.5v569h96q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5t-22.5 9.5zm-359-128l-23-64h-207l115 64h115zm-347-64l-23 64h140l-115-64h-2zm23-64h241l-201-112zm258-55l-26-73h-106zm-119-332l-26 71l65 39zm-44 124l-26 71h140l-1-4zm428-156l-335-33l207 580h128V349z"/>`),
		g.Group(children),
	)
}