package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneThousandOneHundredTwelveDelivery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.447 29.428c.666-2.243 3.31-4.065 5.567-3.645c1.459.28 2.304 1.682 1.986 3.224c-.192 1.122-.987 2.383-2.036 3.084c-1.916 1.262-7.764 4.906-7.764 4.906h7.43M17.328 27.25l3.286-1.532L16.933 37m10.113-24.468L30.332 11l-3.681 11.282m-8.616-9.75L21.32 11l-3.68 11.282"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}