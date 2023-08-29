package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M367.1 160c.6-5.3.9-10.6.9-16C368 64.5 303.5 0 224 0S80 64.5 80 144c0 5.4.3 10.7.9 16H80c-26.5 0-48 21.5-48 48s21.5 48 48 48h288c26.5 0 48-21.5 48-48s-21.5-48-48-48h-.9zM96 288l104.8 209.7c4.4 8.8 13.3 14.3 23.2 14.3s18.8-5.5 23.2-14.3L352 288H96z"/>`),
		g.Group(children),
	)
}