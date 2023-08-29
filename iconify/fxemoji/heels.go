package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heels(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FF473E" d="M11.8 7.3c-.3-.4-.9-.6-1.4-.3C4.2 10.5 0 17.2 0 24.8c0 3.1.7 6.1 1.9 8.7c.5 1 5.1 10.4 5.1 28.2c0 1 4.2 1 4.2 0c0 0 .1-9.7.1-16.2s1.9-8.4 3.5-8.4s8.3 13.2 12 17.7s8.8 8.8 26 8.8s18.9-3.9 18.9-5.4s-2.2-5.6-18-9.8c-16.8-4.4-21-8.2-30.2-22.8C16.4 14.4 13 9.1 11.8 7.3z"/><path fill="#D32A2A" d="M25.2 32.2L12.6 8.4c1.2 1.8 4.1 6 11.1 17.2c8.6 13.6 12.8 17.8 27.1 21.9c4.8 1.4 4 3.6-6.2 2.9c-9.2-.6-15.8-11.3-19.4-18.2zM9 60.9c-.7 0-1.3-.1-1.9-.2v1.1c0 1 4.2 1 4.2 0v-1.1c-.8.1-1.5.2-2.3.2z"/>`),
		g.Group(children),
	)
}