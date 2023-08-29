package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scrutinizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 170.642C0 93.151 37.73 0 153.23 0h293.512c-1.394 75.84-72.898 119.27-132.869 120.718H146.93c-52.324 0-38.303 49.924-38.303 49.924H0zm0 26.561c0 37.326 44.102 118.585 142.112 118.585h171.767c33.473 0 34.794 74.653 1.322 74.653H142.11C85.958 390.441 4.038 429.42.735 512H313.88c69.929 0 130.148-73.332 132.13-152.609c2.023-80.948-53.546-162.188-122.88-162.188H0z"/>`),
		g.Group(children),
	)
}