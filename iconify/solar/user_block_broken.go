package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="6" r="4"/><path stroke-linecap="round" d="M15.5 13.535C14.47 13.195 13.275 13 12 13c-3.866 0-7 1.79-7 4c0 .345 0 .68.027 1M13 20.987c-.316.009-.65.013-1 .013c-1.722 0-3.02-.108-4-.305m11.95-3.645l-3.9 3.9"/><circle cx="18" cy="19" r="3"/></g>`),
		g.Group(children),
	)
}