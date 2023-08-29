package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><ellipse cx="24" cy="11" rx="20" ry="6"/><path fill="#2F88FF" d="M10.7709 15.5C6.61878 16.5994 4 18.208 4 20C4 23.3137 12.9543 26 24 26C35.0457 26 44 23.3137 44 20C44 18.208 41.3812 16.5994 37.2291 15.5C33.7035 16.4335 29.0722 17 24 17C18.9278 17 14.2965 16.4335 10.7709 15.5Z"/><path d="M10.7709 24.5C6.61878 25.5994 4 27.208 4 29C4 32.3137 12.9543 35 24 35C35.0457 35 44 32.3137 44 29C44 27.208 41.3812 25.5994 37.2291 24.5"/><path d="M10.7709 33.5C6.61878 34.5994 4 36.208 4 38C4 41.3137 12.9543 44 24 44C35.0457 44 44 41.3137 44 38C44 36.208 41.3812 34.5994 37.2291 33.5"/></g>`),
		g.Group(children),
	)
}