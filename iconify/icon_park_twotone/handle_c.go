package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandleC0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><path d="M32 17.618c-.898-1.83-3.593-5.031-8.983-4.574c-5.39.458-9.433 5.49-8.983 11.893c.45 6.404 5.39 10.063 9.881 10.063C29.305 35 32 30.609 32 30.609"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandleC0)"/>`),
		g.Group(children),
	)
}