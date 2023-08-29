package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#00c4b3" d="m35.2 34.9l-8.3-8.3v59.7l.1 2.8c0 1.3.2 2.8.7 4.3l65.6 23.1l16.3-7.2l-74.4-74.4z"/><path fill="#22d3c5" d="M27.7 93.4zm81.9 15.9l-16.3 7.2l-65.4-23.1c1.3 4.8 4 10.1 7 13.2l21.3 21.2l47.6.1l5.8-18.6z"/><path fill="#0075c9" d="M1.7 65.1C-.4 67.3.7 72 4 75.5l14.7 14.8l9.2 3.3c-.3-1.5-.7-3-.7-4.3l-.1-2.8l-.2-59.8m82.7 82.6l7.2-16.4l-23-65.6c-1.5-.3-3-.6-4.3-.7l-2.9-.1l-59.6.1"/><path fill="#00a8e1" d="M93.6 27.3c.2 0 .2 0 0 0c.2 0 .2 0 0 0zm16 82l17.7-5.8V54.8l-20.4-20.5c-3-3-8.3-5.8-13.2-7l23.1 65.6"/><path fill="#00c4b3" d="M90.5 18.2L75.7 3.5c-3.4-3.4-8-4.4-10.4-2.3L26.9 26.6h59.5l2.9.1c1.3 0 2.8.2 4.3.7l-3.1-9.2z"/>`),
		g.Group(children),
	)
}