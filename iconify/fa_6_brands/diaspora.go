package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diaspora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M251.64 354.55c-1.4 0-88 119.9-88.7 119.9S76.34 414 76 413.25s86.6-125.7 86.6-127.4c0-2.2-129.6-44-137.6-47.1c-1.3-.5 31.4-101.8 31.7-102.1c.6-.7 144.4 47 145.5 47c.4 0 .9-.6 1-1.3c.4-2 1-148.6 1.7-149.6c.8-1.2 104.5-.7 105.1-.3c1.5 1 3.5 156.1 6.1 156.1c1.4 0 138.7-47 139.3-46.3c.8.9 31.9 102.2 31.5 102.6c-.9.9-140.2 47.1-140.6 48.8c-.3 1.4 82.8 122.1 82.5 122.9s-85.5 63.5-86.3 63.5c-1-.2-89-125.5-90.9-125.5z"/>`),
		g.Group(children),
	)
}