package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchSquareLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 12c0-2.809 0-4.213.674-5.222a4 4 0 0 1 1.104-1.104C7.787 5 9.19 5 12 5c2.809 0 4.213 0 5.222.674a4 4 0 0 1 1.104 1.104C19 7.787 19 9.19 19 12c0 2.809 0 4.213-.674 5.222a4.003 4.003 0 0 1-1.104 1.104C16.213 19 14.81 19 12 19c-2.809 0-4.213 0-5.222-.674a4.002 4.002 0 0 1-1.104-1.104C5 16.213 5 14.81 5 12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.077L14 14" opacity=".5"/><path d="m16.778 5.5l-.082-.368c-.334-1.501-.5-2.252-1.049-2.692C15.1 2 14.33 2 12.791 2H11.21c-1.54 0-2.31 0-2.857.44c-.549.44-.715 1.19-1.05 2.692l-.08.368m9.555 13l-.082.368c-.334 1.501-.5 2.252-1.049 2.692c-.548.44-1.318.44-2.856.44H11.21c-1.539 0-2.308 0-2.856-.44c-.549-.44-.715-1.19-1.05-2.692l-.08-.368" opacity=".5"/></g>`),
		g.Group(children),
	)
}