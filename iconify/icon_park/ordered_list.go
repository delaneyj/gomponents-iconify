package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 4V13"/><path d="M12 13H6"/><path d="M12 27H6"/><path d="M6 19.9998C6 19.9998 9 16.9998 11 19.9998C13 22.9999 6 26.9998 6 26.9998"/><path d="M6.00016 34.5001C6.00016 34.5001 8.00016 31.5 11.0002 33.5C14.0002 35.5 11.0002 38 11.0002 38C11.0002 38 14.0002 40.5 11.0002 42.5C8.00015 44.5 6.00015 41.5 6.00015 41.5"/><path d="M11 38H9"/><path d="M9 4L6 6"/><path d="M21 24H43"/><path d="M21 38H43"/><path d="M21 10H43"/></g>`),
		g.Group(children),
	)
}