package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalfCookie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.8 14c-.927 4.564-4.962 8-9.8 8c-5.523 0-10-4.477-10-10c0-5.185 3.947-9.449 9-9.95"/><path d="M6.5 10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm14-6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-5-3.99l.01-.011m9.99.011l.01-.011M11 12.01l.01-.011M21 9.01l.01-.011M17 6.01l.01-.011M11 2c-.5 1.5.5 3 2.085 3C11 8.5 13 12 18 11.5c0 2.5 2.5 3 3.7 2.514"/></g>`),
		g.Group(children),
	)
}