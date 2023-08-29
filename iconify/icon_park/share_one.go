package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M35 16C37.7614 16 40 13.7614 40 11C40 8.23858 37.7614 6 35 6C32.2386 6 30 8.23858 30 11C30 13.7614 32.2386 16 35 16Z"/><path fill="#2F88FF" d="M13 29C15.7614 29 18 26.7614 18 24C18 21.2386 15.7614 19 13 19C10.2386 19 8 21.2386 8 24C8 26.7614 10.2386 29 13 29Z"/><path stroke-linecap="round" d="M30.0004 13.5745L17.3393 21.2454"/><path stroke-linecap="round" d="M17.3385 26.5639L30.6789 34.4469"/><path fill="#2F88FF" d="M35 32C37.7614 32 40 34.2386 40 37C40 39.7614 37.7614 42 35 42C32.2386 42 30 39.7614 30 37C30 34.2386 32.2386 32 35 32Z"/></g>`),
		g.Group(children),
	)
}