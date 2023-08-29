package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M18 6.017c1.553.047 2.48.22 3.121.862C22 7.757 22 9.172 22 12c0 2.828 0 4.243-.879 5.121C20.243 18 18.828 18 16 18H8c-2.828 0-4.243 0-5.121-.879C2 16.243 2 14.828 2 12c0-2.828 0-4.243.879-5.121C3.757 6 5.172 6 8 6h2"/><path fill="currentColor" d="M18.33 20.335a.75.75 0 1 0 1.34-.67l-1.34.67Zm-1-2l1 2l1.34-.67l-1-2l-1.34.67Zm-11.66 2a.75.75 0 1 1-1.34-.67l1.34.67Zm1-2l-1 2l-1.34-.67l1-2l1.34.67Z" opacity=".6"/><circle cx="14" cy="9" r="5" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 9a2 2 0 1 0 2-2" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.5 9.5V11"/></g>`),
		g.Group(children),
	)
}