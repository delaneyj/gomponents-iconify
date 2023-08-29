package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SatelliteLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g opacity=".5"><path stroke="currentColor" stroke-width="1.5" d="M20.47 10.918s-1.847-.615-4.31-3.078c-2.462-2.463-3.078-4.31-3.078-4.31"/><path fill="currentColor" d="M16.69 8.37a.75.75 0 0 0-1.06-1.06l1.06 1.06Zm-15.054.661a.75.75 0 0 0 .728 1.312L1.636 9.03Zm12.022 12.605a.75.75 0 0 0 1.31.728l-1.31-.728ZM4.47 18.47a.75.75 0 1 0 1.06 1.06l-1.06-1.06Zm8.248-15.595L1.636 9.03l.728 1.312l11.082-6.157l-.728-1.311Zm7.096 7.679l-6.156 11.082l1.31.728l6.157-11.082l-1.31-.728ZM15.63 7.31L4.47 18.47l1.06 1.06L16.69 8.37l-1.06-1.06Z"/></g><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13.082 10.918A5.224 5.224 0 1 0 20.47 3.53a5.224 5.224 0 0 0-7.388 7.388Z"/></g>`),
		g.Group(children),
	)
}