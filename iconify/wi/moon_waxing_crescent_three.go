package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaxingCrescentThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M15.01 25.71c2.04 0 3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09s-3.61-1.51-5.65-1.51c1.71 1.26 2.97 2.9 3.78 4.91S20 12.24 20 14.44c0 .9-.03 1.73-.1 2.5s-.21 1.59-.43 2.47s-.51 1.68-.86 2.4s-.83 1.42-1.45 2.12s-1.33 1.28-2.15 1.78z"/>`),
		g.Group(children),
	)
}