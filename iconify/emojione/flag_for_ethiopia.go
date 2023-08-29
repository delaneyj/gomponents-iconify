package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForEthiopia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M32 62c13.1 0 24.2-8.3 28.3-20H3.7C7.8 53.7 18.9 62 32 62z"/><path fill="#699635" d="M32 2C18.9 2 7.8 10.4 3.7 22h56.6C56.2 10.4 45.1 2 32 2z"/><path fill="#ffce31" d="M60.3 42c1.1-3.1 1.7-6.5 1.7-10s-.6-6.9-1.7-10H3.7C2.6 25.1 2 28.5 2 32s.6 6.9 1.7 10h56.6"/><circle cx="32" cy="32" r="14" fill="#428bc1"/><g fill="#ffce31"><path d="m35.8 33.3l6.1-4.3h-7.5L32 22l-2.3 7h-7.5l6.1 4.3l-2.3 7l6-4.3l6.1 4.3l-2.3-7m3.5-3.5l-3.8 2.7l-.9-2.7h4.7M29.2 33l1.1-3.2h3.5l1.1 3.2l-2.9 2l-2.8-2m2.8-8.4l1.5 4.4h-2.9l1.4-4.4m-7.3 5.2h4.7l-.9 2.7l-3.8-2.7m2.8 8.4l1.4-4.4l2.4 1.7l-3.8 2.7m9 0l-3.8-2.7l2.4-1.7l1.4 4.4"/><path d="M31.6 37.3V42h.8v-4.7L32 37zm-5.7-4.6L22 33.8l.2.8l4.6-1.3zm3.2-4.6l.2-.8l-3.1-3.8l-.6.5l3.3 4.1zm6 0l3.3-4.1l-.6-.5l-3.1 3.8l.2.8zm3 4.6l-.9.6l4.6 1.3l.2-.8z"/></g>`),
		g.Group(children),
	)
}