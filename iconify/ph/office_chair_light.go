package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeChairLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246 128a6 6 0 0 1-6 6h-18.39A46.07 46.07 0 0 1 176 174h-42v28h26a30 30 0 0 1 30 30a6 6 0 0 1-12 0a18 18 0 0 0-18-18h-26v18a6 6 0 0 1-12 0v-18H96a18 18 0 0 0-18 18a6 6 0 0 1-12 0a30 30 0 0 1 30-30h26v-28H80a46.07 46.07 0 0 1-45.61-40H16a6 6 0 0 1 0-12h24a6 6 0 0 1 6 6a34 34 0 0 0 34 34h96a34 34 0 0 0 34-34a6 6 0 0 1 6-6h24a6 6 0 0 1 6 6Zm-176.57 9.17A14 14 0 0 1 66.14 126l13.72-96a14.07 14.07 0 0 1 13.86-12h68.56a14.07 14.07 0 0 1 13.86 12l13.72 96A14 14 0 0 1 176 142H80a14 14 0 0 1-10.57-4.83Zm9.06-7.86A2 2 0 0 0 80 130h96a2 2 0 0 0 1.51-.69a2 2 0 0 0 .47-1.59l-13.72-96a2 2 0 0 0-2-1.72H93.72a2 2 0 0 0-2 1.72l-13.72 96a2 2 0 0 0 .49 1.59Z"/>`),
		g.Group(children),
	)
}