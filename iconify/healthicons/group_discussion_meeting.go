package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupDiscussionMeeting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M24 12.042A6.042 6.042 0 0 0 17.958 6h-5.916a6.042 6.042 0 1 0 0 12.083H13.2V21S24 19.542 24 12.042ZM18.5 28c0 2.21-1.79 4-4 4s-4-1.79-4-4s1.79-4 4-4s4 1.79 4 4Zm-4 6C11.663 34 6 35.43 6 38.267V42h17v-3.733C23 35.429 17.337 34 14.5 34Zm19-2c2.21 0 4-1.79 4-4s-1.79-4-4-4s-4 1.79-4 4s1.79 4 4 4Zm0 2c-2.837 0-8.5 1.43-8.5 4.267V42h17v-3.733C42 35.429 36.337 34 33.5 34ZM32.042 7A6.042 6.042 0 0 0 26 13.042c0 7.5 9.6 8.958 9.6 8.958v-2.917h.358a6.042 6.042 0 1 0 0-12.083h-3.916Z"/>`),
		g.Group(children),
	)
}