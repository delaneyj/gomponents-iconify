package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JustTwentyFourHours(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.21 32.84V14.76l-6.75 8.56l-.72.84l-2.17 2.77H37.5m-27-6.14a6 6 0 0 1 6-6a6 6 0 0 1 6 6A5.33 5.33 0 0 1 20.75 25c-2.42 2.16-10.25 7.83-10.25 7.83h11.93"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}