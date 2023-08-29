package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M46.25 45.7c0 1-10 1.3-22.25 1.3s-22.25-.31-22.25-1.3s10-1.81 22.25-1.81s22.25.81 22.25 1.81Z" opacity=".15"/><path fill="#fff" d="M3.03 3.03h41.94v41.94H3.03Z"/><path fill="#f0f0f0" d="M41.37 42H6.63A3.6 3.6 0 0 1 3 38.38v3A3.6 3.6 0 0 0 6.63 45h34.74a3.6 3.6 0 0 0 3.6-3.6v-3a3.6 3.6 0 0 1-3.6 3.6Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.03 3.03h41.94v41.94H3.03Z"/><path fill="#bf8df2" d="M6.03 6.03h35.95v35.95H6.03Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M6.03 6.03h35.95v35.95H6.03Z"/><path fill="#6dd627" d="M6 26.41h36v14.36a1.2 1.2 0 0 1-1.2 1.2H7.23A1.2 1.2 0 0 1 6 40.77V26.41Z"/><path fill="#46b000" d="M40.77 39H7.23A1.2 1.2 0 0 1 6 37.78v3A1.19 1.19 0 0 0 7.23 42h33.54a1.19 1.19 0 0 0 1.2-1.2v-3a1.2 1.2 0 0 1-1.2 1.2Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.91 17.63a3.22 3.22 0 1 1 0-6.44a3.11 3.11 0 0 1 .43 0a4.81 4.81 0 0 0-3.05-1.08a4.87 4.87 0 1 0 4.4 7a3.22 3.22 0 0 1-1.78.52Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M6 26.41h36v14.36a1.2 1.2 0 0 1-1.2 1.2H7.23A1.2 1.2 0 0 1 6 40.77V26.41h0Z"/><path fill="#ffaa54" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 22.29L21.38 36.6l8.43-3.77a.59.59 0 0 0 .26-.86Z"/><path fill="#ff8a14" d="m24 22.29l-10.25-2.4a.6.6 0 0 0-.72.46c-.15.76-.51 2.35-1.31 5.39a34 34 0 0 1-2.54 6.84a.59.59 0 0 0 .38.84l11.82 3.18Z"/><path fill="#ffaa54" d="m23.27 26.28l.73-4l-10.24-2.4a.6.6 0 0 0-.73.46c-.12.61-.38 1.76-.89 3.77Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m24 22.29l-10.25-2.4a.6.6 0 0 0-.72.46c-.15.76-.51 2.35-1.31 5.39a34 34 0 0 1-2.54 6.84a.59.59 0 0 0 .38.84l11.82 3.18Zm0 0l2.07 12.21"/>`),
		g.Group(children),
	)
}