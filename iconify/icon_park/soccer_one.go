package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoccerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-miterlimit="2" stroke-width="4" d="M29 14C31.7614 14 34 11.7614 34 9C34 6.23858 31.7614 4 29 4C26.2386 4 24 6.23858 24 9C24 11.7614 26.2386 14 29 14Z"/><path fill="#000" d="M19 46C20.6569 46 22 44.6569 22 43C22 41.3431 20.6569 40 19 40C17.3431 40 16 41.3431 16 43C16 44.6569 17.3431 46 19 46Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M24.38 18.9102L20 27.0002L27.31 33.0202C27.62 33.2602 27.84 33.5902 27.97 33.9602L31 44.0002"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M20 27L14.51 34.64C13.97 35.43 12.95 35.73 12.07 35.36L4 32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M42 23.9999H35C34.64 23.9999 34.28 23.9099 33.96 23.7299L30.75 21.9199C26.74 19.5699 22.35 17.9599 17.78 17.1399L14.56 16.5699C14.11 16.4899 13.64 16.5699 13.24 16.7899L7 19.9999"/></g>`),
		g.Group(children),
	)
}