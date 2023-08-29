package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeCtaskmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M43.5 28.8v.2c0 5.4-4.4 9.8-9.8 9.8h0c-5.4 0-9.8-4.4-9.8-9.8V19c0-5.4 4.4-9.8 9.8-9.8h0c5.4 0 9.8 4.4 9.8 9.8v.1"/><path d="M38.9 24.7c.1-.5.1-.9 0-1.4l1.5-1.2c.1-.1.2-.3.1-.4l-1.4-2.4c-.1-.1-.3-.2-.4-.2l-1.7.7c-.4-.3-.8-.5-1.2-.7l-.3-1.9c0-.2-.2-.3-.3-.3h-2.8c-.2 0-.3.1-.3.3l-.3 1.9c-.4.2-.8.4-1.2.7l-1.7-.7c-.2-.1-.3 0-.4.2L27 21.7c-.1.1-.1.3.1.4l1.5 1.2c-.1.5-.1.9 0 1.4l-1.5 1.2c-.1.1-.2.3-.1.4l1.4 2.4c.1.1.3.2.4.2l1.7-.7c.4.3.8.5 1.2.7l.3 1.9c0 .2.2.3.3.3h2.8c.2 0 .3-.1.3-.3l.3-1.9c.4-.2.8-.4 1.2-.7l1.7.7c.2.1.3 0 .4-.2l1.4-2.4c.1-.1.1-.3-.1-.4l-1.4-1.2zM33.7 26c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.8 2-2 2c.1 0 0 0 0 0z"/><g stroke-miterlimit="10"><path d="M16.6 24c4 0 7.4 3.3 7.4 7.4h0c0 4-3.3 7.4-7.4 7.4h-2.9c-5.1 0-7-.7-9.2-2.6"/><path d="M4.5 11.9c2.2-1.8 4.4-2.6 9.2-2.6h2.9c4 0 7.4 3.3 7.4 7.4h0c0 4-3.3 7.4-7.4 7.4H9.3"/></g></g>`),
		g.Group(children),
	)
}