package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinOffTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l7.572 7.573l-3.625 1.394a1.75 1.75 0 0 0-.61 2.871L8.69 18.25L3 23.939v1.06h1.06l5.69-5.689l3.132 3.132a1.75 1.75 0 0 0 2.87-.61l1.395-3.625l7.572 7.573a.75.75 0 0 0 1.061-1.06L3.28 2.22Zm8.52 6.397l7.584 7.584l4.232-2.005c2.031-.962 2.496-3.646.907-5.235L19.04 3.478c-1.59-1.59-4.274-1.125-5.236.907L11.8 8.617Z"/>`),
		g.Group(children),
	)
}