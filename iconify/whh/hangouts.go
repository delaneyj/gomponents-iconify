package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hangouts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 1024V830q-106-8-194-65T51 616T0 416q0-113 55.5-209T207 55.5T416 0t209 55.5T776.5 207T832 416q0 67-15.5 130.5T775 661t-59.5 97.5T645 841t-74 66.5t-71 51.5t-59.5 36t-41.5 22zm0-730q0-16-11-27t-26-11H229q-15 0-26 11t-11 27v122q0 32 32 32h96q0 64-64 64v64q51 0 89.5-34t38.5-88V294zm256 0q0-16-11-27t-26-11H485q-15 0-26 11t-11 27v122q0 32 32 32h96q0 64-64 64v64q51 0 89.5-34t38.5-88V294z"/>`),
		g.Group(children),
	)
}