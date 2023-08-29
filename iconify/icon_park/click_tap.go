package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClickTap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M15 26V15C15 13.3431 16.3431 12 18 12C19.6569 12 21 13.3431 21 15V26"/><path stroke-linejoin="round" d="M39 25V31.5C39 37.8513 33.8513 43 27.5 43H26.5C20.1487 43 15 37.8513 15 31.5V25"/><path stroke-linejoin="round" d="M21 29V27.1058V24C21 22.3431 22.3431 21 24 21C25.6569 21 27 22.3431 27 24V27.1818V29"/><path stroke-linejoin="round" d="M27 29V27.1058V24C27 22.3431 28.3431 21 30 21C31.6569 21 33 22.3431 33 24V27.1818V29"/><path stroke-linejoin="round" d="M33 29V27.1058V24C33 22.3431 34.3431 21 36 21C37.6569 21 39 22.3431 39 24V27.1818V29"/><path d="M28 15C28 13.9489 27.8378 12.9357 27.5371 11.9839C27.2008 10.9195 26.6913 9.93208 26.0415 9.0547C24.2198 6.59472 21.2961 5 18 5C14.7039 5 11.7802 6.59472 9.95847 9.0547C9.30873 9.93208 8.79916 10.9195 8.46286 11.9839C8.16217 12.9357 8 13.9489 8 15"/></g>`),
		g.Group(children),
	)
}