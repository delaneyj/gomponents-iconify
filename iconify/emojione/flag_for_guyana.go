package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGuyana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#699635" d="M32 2c-5.3 0-10.3 1.4-14.6 3.8l-.1 52.3C21.6 60.6 26.6 62 32 62c14.9 0 27.3-10.9 29.6-25.2v-9.6C59.3 12.9 46.9 2 32 2z"/><path fill="#ffe62e" d="M60.1 29.6L17.4 8v.4h-3.8c-.3.3-.7.5-1 .8v45.5c.3.3.7.6 1 .8h3.7v.5l43.2-21.5c.1-.8 1.6-1.7 1.6-2.5s-2-1.6-2-2.4"/><path fill="#fff" d="M13.5 8.4L62 31.8c0-1.6-.1-3.1-.4-4.6L17.4 5.8c-1.4.8-2.7 1.6-3.9 2.6"/><path fill="#ed4c5c" d="M12.5 12H9.6C4.9 17.3 2 24.3 2 32s2.9 14.7 7.6 20h2.8l20-20l-19.9-20z"/><path fill="#3e4347" d="M12.5 9.2c-1 .9-2 1.8-2.9 2.8l20 20l-20 20c.9 1 1.8 1.9 2.9 2.8L35.3 32L12.5 9.2"/><path fill="#fff" d="M61.6 36.8c.2-1.5.4-3.1.4-4.6L13.5 55.6c1.2 1 2.5 1.8 3.9 2.6l44.2-21.4"/>`),
		g.Group(children),
	)
}