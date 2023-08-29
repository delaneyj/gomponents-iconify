package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H38C39.1046 44 40 43.1046 40 42V6C40 4.89543 39.1046 4 38 4Z"/><path fill="#2F88FF" d="M28 16C28 17.3807 27.4404 18.6307 26.5355 19.5355C25.6307 20.4404 24.3807 21 23 21C20.2386 21 18 18.7614 18 16C18 13.2386 20.2386 11 23 11C25.7614 11 28 13.2386 28 16Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 23L26.5355 19.5355M26.5355 19.5355C27.4404 18.6307 28 17.3807 28 16C28 13.2386 25.7614 11 23 11C20.2386 11 18 13.2386 18 16C18 18.7614 20.2386 21 23 21C24.3807 21 25.6307 20.4404 26.5355 19.5355Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 30L31 30"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 36H24"/></g>`),
		g.Group(children),
	)
}