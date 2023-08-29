package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffce31" d="M2 28.3v31.4C2 62.1 3.9 64 6.3 64h51.4c2.4 0 4.3-1.9 4.3-4.3V28.3H2z"/><path fill="#ff8736" d="M62 24c0-2.4-1.9-4.3-4.3-4.3H6.3C3.9 19.6 2 21.6 2 24v4.3h60V24z"/><g fill="#3e4347"><ellipse cx="12.4" cy="23.5" rx="5.9" ry="2.5"/><ellipse cx="51.6" cy="23.5" rx="5.9" ry="2.5"/><path d="m36.6 56.4l-1.9-12.3c1.1-.8 1.9-2.2 1.9-3.7c0-2.5-2-4.6-4.6-4.6s-4.6 2.1-4.6 4.6c0 1.5.7 2.9 1.9 3.7l-1.9 12.3h9.2"/></g><path fill="#dfe9ef" d="M32 0c-5.6 0-10.7 2-14.8 5.3l2.9 2.9c3.3-2.7 7.4-4.3 11.9-4.3c10.4 0 18.9 8.7 19.6 19.7V25c2.2 0 3.8-.6 3.8-1.4C55.4 10.6 44.9 0 32 0"/><path fill="#b0bdc6" d="M51.6 23.5C50.9 12.6 42.4 3.9 32 3.9c-4.4 0-8.6 1.6-11.9 4.3L23 11c2.4-1.8 5.5-3 9-3c9.5 0 15.5 8.4 15.5 15.5c0 .8 2 1.4 4.2 1.4l-.1-1.4"/>`),
		g.Group(children),
	)
}