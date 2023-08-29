package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 22c-4.418 0-8-3.646-8-8.143c0-4.462 2.553-9.67 6.537-11.531A3.453 3.453 0 0 1 12 2v20Z" clip-rule="evenodd"/><path d="M13.463 2.326A3.453 3.453 0 0 0 12 2v7l4.432-4.432c-.863-.947-1.86-1.724-2.97-2.242Z" opacity=".3"/><path d="M12 9v5.5l6.614-6.614c-.572-1.22-1.308-2.357-2.182-3.318L12 9Z" opacity=".4"/><path d="m12 19.5l7.811-7.811a14.94 14.94 0 0 0-1.197-3.803L12 14.5v5Z" opacity=".6"/><path d="M19.811 11.689L12 19.5V22c4.418 0 8-3.646 8-8.143c0-.71-.064-1.438-.189-2.168Z" opacity=".7"/></g>`),
		g.Group(children),
	)
}