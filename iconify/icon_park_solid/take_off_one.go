package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TakeOffOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTakeOffOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4.997 40.988h38"/><path fill="#fff" d="m8.52 31.264l-4.612-7.99c.97-.56 5.771 1.115 7.559 2.032l9.702-3.473l-8.296-14.368l4.115-.247l13.4 12.462l8.25-2.612c3.655-1.045 4.807.95 5.037 1.35c1.383 2.394-1.411 4.007-1.81 4.238c-3.193 1.843-33.344 8.608-33.344 8.608Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTakeOffOne0)"/>`),
		g.Group(children),
	)
}