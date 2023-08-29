package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 768q0-156 61.5-298.5T226 224T471 60.5T769 0t298.5 60.5T1313 224t164 245t61 299t-61 299t-164 244.5t-245.5 163T769 1535t-298.5-61T225 1310T61 1065T0 768zm170 0q0 245 177 422q176 176 422 176q163 0 301.5-80.5t219-218T1370 768q0-121-47.5-232T1194 344.5t-192-128T769 169q-121 0-231.5 47.5t-191 128T218 536t-48 232zm334-259q-5-11 1-16.5t16-.5l238 89q10 4 23 0l235-89q10-5 16 .5t2 16.5l-253 599q-3 10-13 10q-7 0-10-10z" fill="currentColor"/>`),
		g.Group(children),
	)
}