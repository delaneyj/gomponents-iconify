package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorRightBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 21c4.243 0 6.364 0 7.682-1.318C21 18.364 21 16.242 21 12m-9-9c4.243 0 6.364 0 7.682 1.318c.816.816 1.127 1.939 1.245 3.682"/><path stroke-dasharray="2.5 3" d="M13 21H9c-2.828 0-4.243 0-5.121-.879C3 19.243 3 17.828 3 15V9c0-2.828 0-4.243.879-5.121C4.757 3 6.172 3 9 3h4"/><path d="M12 22V2"/></g>`),
		g.Group(children),
	)
}