package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRemoveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFileRemoveOutline0"><g id="evaFileRemoveOutline1"><g id="evaFileRemoveOutline2" fill="currentColor"><path d="m19.74 8.33l-5.44-6a1 1 0 0 0-.74-.33h-7A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V9a1 1 0 0 0-.26-.67ZM14 5l2.74 3h-2a.79.79 0 0 1-.74-.85Zm3.44 15H6.56a.53.53 0 0 1-.56-.5v-15a.53.53 0 0 1 .56-.5H12v3.15A2.79 2.79 0 0 0 14.71 10H18v9.5a.53.53 0 0 1-.56.5Z"/><path d="M14 13h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/></g></g></g>`),
		g.Group(children),
	)
}