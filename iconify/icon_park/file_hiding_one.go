package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHidingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23V14L31 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H22"/><path d="M25 30C25.2821 30.9145 25.7095 31.7628 26.2546 32.5166C27.7827 34.63 30.2351 36 33 36C35.7649 36 38.2173 34.63 39.7454 32.5166C40.2905 31.7628 40.7179 30.9145 41 30"/><path d="M30.5215 36.0684L29.4862 39.9321"/><path d="M35.4863 36.0684L36.5216 39.9321"/><path d="M39.3535 34.3536L42.1819 37.182"/><path d="M24 37.0103L26.8285 34.1819"/><path d="M30 4V14H40"/></g>`),
		g.Group(children),
	)
}