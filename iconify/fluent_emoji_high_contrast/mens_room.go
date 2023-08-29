package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MensRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.941 10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm3.332.938h-6.546c-1.23 0-2.227.997-2.227 2.227v5.156c0 .539.414 1.002.952 1.027a1 1 0 0 0 1.048-.999v-4.372a.501.501 0 1 1 1.002 0v11.995c0 .539.414 1.002.952 1.027A1 1 0 0 0 15.502 26v-5.852a.5.5 0 1 1 1 0v5.824c0 .539.414 1.002.952 1.027A1 1 0 0 0 18.503 26V13.975a.499.499 0 0 1 .997 0v4.375a1 1 0 0 0 1.048.999c.538-.026.952-.489.952-1.027v-5.156a2.228 2.228 0 0 0-2.227-2.229Z"/><path fill-rule="evenodd" d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}