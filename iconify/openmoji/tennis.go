package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M43.7 63.9c.3-4.4 1.9-6.6 5.4-10.1c3.3-3.3 5.5-7.4 5.5-12.2c0-9.5-8.3-17.2-18.6-17.2s-18.6 7.7-18.6 17.2c0 4.8 2.2 8.9 5.5 12.2c3.5 3.5 5.1 5.7 5.4 10.1"/><circle cx="36" cy="36" r="29" fill="#FCEA2B"/><path fill="#F1B31C" d="M50.7 11c4.8 5.2 7.8 12.1 7.8 19.8c0 16-13 29-29 29c-5.4 0-10.4-1.5-14.7-4C20.1 61.4 27.7 65 36 65c16 0 29-13 29-29c0-10.7-5.7-20-14.3-25z"/><path fill="none" stroke="#FFF" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M43.7 63.9c.3-4.4 1.9-6.6 5.4-10.1c3.3-3.3 5.5-7.4 5.5-12.2c0-9.5-8.3-17.2-18.6-17.2s-18.6 7.7-18.6 17.2c0 4.8 2.2 8.9 5.5 12.2c3.5 3.5 5.1 5.7 5.4 10.1"/><circle cx="36" cy="36" r="29" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}