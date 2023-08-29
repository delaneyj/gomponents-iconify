package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928.784 1024q-40 0-68-28t-28-68q0-150-58.5-286t-157-234.5t-234.5-157t-286-58.5q-40 0-68-28t-28-68t28-68t68-28q126 0 246.5 33.5t222 93t187.5 145.5t145.5 187.5t93 222t33.5 246.5q0 40-28 68t-68 28zm-832-640q111 0 211.5 43t173.5 116t116 173.5t43 211.5q0 40-28 68t-68 28t-68-28t-28-68q0-151-100.5-251.5T96.784 576q-40 0-68-28t-28-68t28-68t68-28zm32 384q53 0 90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5t-90.5-37.5T.784 896t37.5-90.5t90.5-37.5z"/>`),
		g.Group(children),
	)
}