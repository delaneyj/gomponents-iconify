package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinterestLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214 112c0 22.05-7.7 42.19-21.68 56.73c-13 13.52-30.61 21.27-48.32 21.27c-19.42 0-31.55-7.07-38.63-13.64l-11.53 49A6 6 0 0 1 88 230a5.89 5.89 0 0 1-1.37-.16a6 6 0 0 1-4.47-7.21l32-136a6 6 0 1 1 11.68 2.74L108.75 162c2.8 4.33 12.46 16 35.25 16c28.51 0 58-24.69 58-66a74 74 0 1 0-138.1 37a6 6 0 1 1-10.39 6A86 86 0 1 1 214 112Z"/>`),
		g.Group(children),
	)
}