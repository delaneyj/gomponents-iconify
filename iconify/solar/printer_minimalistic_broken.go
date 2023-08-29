package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 12c0 2.828 0 4.243-.879 5.121c-.641.642-1.567.815-3.121.862m-12 0c-1.553-.047-2.48-.22-3.121-.862C2 16.243 2 14.828 2 12c0-2.828 0-4.243.879-5.121C3.757 6 5.172 6 8 6h8c2.828 0 4.243 0 5.121.879c.3.3.498.662.628 1.121"/><path d="M17.983 6c-.047-1.553-.22-2.48-.862-3.121C16.243 2 14.828 2 12 2c-2.828 0-4.243 0-5.121.879c-.642.641-.815 1.568-.862 3.121M18 12v4c0 2.828 0 4.243-.879 5.121c-.641.642-1.567.815-3.121.862M6 12v4c0 2.828 0 4.243.879 5.121c.641.642 1.568.815 3.121.862"/><path d="M12 11c3.412 0 5.873.713 7.5 1.443m-15 0a15.792 15.792 0 0 1 3.5-1.08"/></g>`),
		g.Group(children),
	)
}