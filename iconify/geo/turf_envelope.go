package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfEnvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="13.162" cy="54.683" r="5.824" fill="currentColor"/><circle cx="51.116" cy="15.612" r="5.824" fill="currentColor"/><circle cx="87.954" cy="83.706" r="5.824" fill="currentColor"/><path fill="currentColor" d="M36.674 82h-9.256a2 2 0 0 0 0 4h9.256a2 2 0 0 0 0-4zm18.513 0H45.93a2 2 0 0 0 0 4h9.256a2 2 0 0 0 .001-4zm18.511 0h-9.256a2 2 0 0 0 0 4h9.256a2 2 0 0 0 0-4zm-55.536 0H15v-3.294a2 2 0 0 0-4 0v5C11 84.811 12.058 86 13.162 86h5a2 2 0 0 0 0-4zM11 64.301v2.786a2 2 0 0 0 4 0v-2.786c0 .134-1.315.206-2 .206s-1-.072-2-.206zm3.708-19.296c.233-.327.292-.724.292-1.156V32.23a2 2 0 0 0-4 0v11.619c0 .432.221.829.454 1.156c.531-.089 1.072-.146 1.627-.146s1.097.057 1.627.146zM13 22.612a2 2 0 0 0 2-2V18h3.162a2 2 0 0 0 0-4h-5C12.058 14 11 14.507 11 15.612v5a2 2 0 0 0 2 2zM27.418 18h9.256a2 2 0 0 0 0-4h-9.256a2 2 0 0 0 0 4zm46.28-4h-9.256a2 2 0 0 0 0 4h9.256a2 2 0 0 0 0-4zm9.256 4H86v2.612a2 2 0 0 0 4 0v-5C90 14.507 89.059 14 87.954 14h-5a2 2 0 0 0 0 4zM90 32.231a2 2 0 0 0-4 0V43.85a2 2 0 0 0 4 0V32.231zm-2 21.238a2 2 0 0 0-2 2v11.618a2 2 0 0 0 4 0V55.469a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}