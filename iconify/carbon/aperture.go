package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aperture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm11.84 12.14L22 17.52V5.62a12 12 0 0 1 5.84 8.52ZM12 18.68v-5.36L16 11l4 2.31v5.36L16 21Zm8-14V11l-9.66-5.58A11.9 11.9 0 0 1 20 4.7Zm-11.52 2L14 9.85L4 15.62a12 12 0 0 1 4.48-8.96ZM4.16 17.85L10 14.47v11.91a12 12 0 0 1-5.84-8.53ZM12 27.3V21l9.67 5.58A11.92 11.92 0 0 1 16 28a12.05 12.05 0 0 1-4-.7Zm11.52-2L18 22.14l10-5.77a12 12 0 0 1-4.48 8.97Z"/>`),
		g.Group(children),
	)
}