package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInBrowserTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm.011 2l-.084.005l-.055.012l-.083.03l-.074.042l-.056.045l-2.513 2.512l-.057.07a.5.5 0 0 0 0 .568l.057.07l.07.057a.5.5 0 0 0 .568 0l.07-.057l1.645-1.646L17 21l.008.09a.5.5 0 0 0 .402.402l.09.008l.09-.008a.5.5 0 0 0 .402-.402L18 21l-.001-5.294l1.647 1.648l.07.057a.5.5 0 0 0 .695-.695l-.057-.07l-2.548-2.542l-.048-.032l-.067-.034l-.063-.021l-.054-.012A.5.5 0 0 0 17.51 14zM6.25 3h11.5a3.25 3.25 0 0 1 3.245 3.066L21 6.25l.001 5.773a6.47 6.47 0 0 0-1.5-.71L19.5 8h-15v9.75a1.75 1.75 0 0 0 1.606 1.744l.144.006h5.064a6.47 6.47 0 0 0 .709 1.501L6.25 21a3.25 3.25 0 0 1-3.245-3.066L3 17.75V6.25a3.25 3.25 0 0 1 3.066-3.245L6.25 3z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}