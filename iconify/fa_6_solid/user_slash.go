package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M38.8 5.1C28.4-3.1 13.3-1.2 5.1 9.2s-6.3 25.5 4.1 33.7l592 464c10.4 8.2 25.5 6.3 33.7-4.1s6.3-25.5-4.1-33.7L353.3 251.6C407.9 237 448 187.2 448 128C448 57.3 390.7 0 320 0c-69.8 0-126.5 55.8-128 125.2L38.8 5.1zm225.5 299.2C170.5 309.4 96 387.2 96 482.3c0 16.4 13.3 29.7 29.7 29.7h388.6c3.9 0 7.6-.7 11-2.1l-261-205.6z"/>`),
		g.Group(children),
	)
}