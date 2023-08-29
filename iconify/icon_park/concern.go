package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Concern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000"><path stroke-linecap="round" stroke-width="4" d="M10.8579 9.85803C7.23858 13.4773 5 18.4773 5 24.0002C5 29.523 7.23858 34.523 10.8579 38.1423"/><path stroke-linecap="round" stroke-width="4" d="M39.1421 38.1423C42.7614 34.523 45 29.523 45 24.0002C45 18.4773 42.7614 13.4773 39.1421 9.85803"/><path stroke-linecap="round" stroke-width="4" d="M34.8994 33.8997C37.4329 31.3662 38.9999 27.8662 38.9999 24.0002C38.9999 20.1342 37.4329 16.6342 34.8994 14.1007"/><path stroke-linecap="round" stroke-width="4" d="M15.1005 14.1007C12.567 16.6342 11 20.1342 11 24.0002C11 27.8662 12.567 31.3662 15.1005 33.8997"/><path fill="#2F88FF" stroke-linejoin="round" stroke-width="3.5" d="M28.1818 20C30.2905 20 32 21.6118 32 23.6C32 26.1882 29.4545 28.4 28.1818 29.6C27.3333 30.4 26.2727 31.2 25 32C23.7273 31.2 22.6667 30.4 21.8182 29.6C20.5455 28.4 18 26.1882 18 23.6C18 21.6118 19.7095 20 21.8182 20C23.1463 20 24.316 20.6393 25 21.6093C25.684 20.6393 26.8537 20 28.1818 20Z"/></g>`),
		g.Group(children),
	)
}