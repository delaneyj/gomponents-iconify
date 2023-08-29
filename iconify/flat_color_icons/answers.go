package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Answers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#42A5F5" d="M36 44H8V8h20l8 8z"/><path fill="#90CAF9" d="M40 40H12V4h20l8 8z"/><path fill="#E1F5FE" d="M38.5 13H31V5.5z"/><path fill="#1976D2" d="M23.4 29.9c0-.2 0-.4.1-.6s.2-.3.3-.5s.3-.2.5-.3s.4-.1.6-.1s.5 0 .7.1s.4.2.5.3s.2.3.3.5s.1.4.1.6s0 .4-.1.6s-.2.3-.3.5s-.3.2-.5.3s-.4.1-.7.1s-.5 0-.6-.1s-.4-.2-.5-.3s-.2-.3-.3-.5s-.1-.4-.1-.6zm2.7-3.1h-2.3l-.4-9.8h3l-.3 9.8z"/>`),
		g.Group(children),
	)
}