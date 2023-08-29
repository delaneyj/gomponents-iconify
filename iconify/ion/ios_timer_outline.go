package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosTimerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M236.6 271.6c4.6 5.7 11.5 9.4 19.4 9.4 13.8 0 25-11.2 25-25 0-7.3-3.2-13.8-8.2-18.4-.6-.7-1.3-1.5-2.2-2.2 0 0-117.7-87.5-120.3-85.2-2.6 2.3 85.3 120.2 85.3 120.2.2.4.7.8 1 1.2z" fill="currentColor"/><path d="M256.2 48h-.2v112h16V65.3c97.8 8.3 175.3 90.5 175.3 190.5 0 105.5-85.7 191.4-191.2 191.4-105.5 0-191.3-85.8-191.3-191.3 0-52.8 21.5-100.6 56.1-135.2L109 108.9C71.3 146.6 48 198.6 48 256c0 114.9 93.1 208 208 208s208-93.1 208-208S371 48 256.2 48z" fill="currentColor"/>`),
		g.Group(children),
	)
}