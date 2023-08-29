package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrimacingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><g fill="#664e27"><path d="M52.5 40c-2.2-3.6-7.7-6-20.5-6s-18.3 2.4-20.5 6c-1.2 1.9.5 5 .5 5c2.1 3.9 1.8 5 20 5c18.1 0 17.9-1.1 20-5c0 0 1.7-3.1.5-5"/><circle cx="20.5" cy="24.5" r="5"/><circle cx="43.5" cy="24.5" r="5"/></g><path fill="#fff" d="M48 40c.1-.5-.2-1.2-.6-1.5c0 0-3.9-2.5-15.4-2.5s-15.4 2.5-15.4 2.5c-.4.3-.7.9-.6 1.5l.2 1c.1.5.6 1 1.1 1h29.3c.5 0 1-.4 1.1-1l.3-1m-16 8c6.3 0 15.2 0 15-2.1c0-.4-.1-.8-.3-1.3c-.2-.5-.3-.7-1.4-.7H18.6c-1.1 0-1.2.1-1.4.7c-.1.5-.2.9-.3 1.3C16.8 48 25.7 48 32 48"/>`),
		g.Group(children),
	)
}