package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Universal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M24 38C31.732 38 38 31.732 38 24C38 16.268 31.732 10 24 10C16.268 10 10 16.268 10 24C10 31.732 16.268 38 24 38Z"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M11 29C12.5089 29.6244 15 30 16.2586 28.5317C17.5171 27.0633 16.395 24.7522 17.7889 23.9682C19.317 23.1086 20.4203 26.0319 23.2912 25.516C26.162 25 28 21 28 19C28 17 26.2854 17 26.162 15.0542C26 12.5 28 11 28 11"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M27.9995 37C26.9137 36.091 25.9995 35.5 26.0001 34C26.0006 32.5 26.9995 33 27.9995 32C28.9996 31 28.4995 29 29.4999 28.5C30.5004 28 33.6075 29.0558 35.9998 31"/><circle cx="24" cy="4" r="2" fill="#000"/><circle cx="24" cy="44" r="2" fill="#000"/><circle cx="44" cy="24" r="2" fill="#000"/><circle cx="38" cy="10" r="2" fill="#000"/><circle cx="10" cy="38" r="2" fill="#000"/><circle cx="4" cy="24" r="2" fill="#000"/><circle cx="10" cy="10" r="2" fill="#000"/><circle cx="38" cy="38" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 24C10 27.8146 11.5256 31.2729 14 33.798"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 38C31.732 38 38 31.732 38 24"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 10C27.8146 10 31.2729 11.5256 33.798 14"/></g>`),
		g.Group(children),
	)
}