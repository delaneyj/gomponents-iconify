package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 .5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4A.5.5 0 0 1 3 .5Zm2 7a.5.5 0 0 0 1 0v-3a.5.5 0 0 0-1 0v3ZM5.5 2a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9ZM2 6.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Zm8.148-2.647a.5.5 0 1 0 .706-.708l-1.002-.998a.5.5 0 1 0-.706.708l1.002.998Z"/>`),
		g.Group(children),
	)
}