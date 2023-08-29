package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13c.6 0 1-.4 1-1V9c0-.6-.4-1-1-1s-1 .4-1 1v3c0 .6.4 1 1 1zm-1 2c0 .5.5 1 1 1s1-.5 1-1c0-.3-.1-.5-.3-.7c-.3-.3-.7-.4-1.1-.2c-.1 0-.2.1-.3.2c-.2.2-.3.4-.3.7zm6-11.7C13.1 1.1 8.3 1.8 5.1 4.7V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .1 0 .2.1.3v.1c.1.2.3.4.5.5c.2.1.3.1.4.1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H6.2C7.7 4.9 9.8 4 12 4c4.4 0 8 3.6 8 8c0 .6.4 1 1 1s1-.4 1-1c0-3.6-1.9-6.9-5-8.7zm2.9 12.2h-4.5c-.6 0-1 .4-1 1s.4 1 1 1h2.4C16.3 19.1 14.2 20 12 20c-4.4 0-8-3.6-8-8c0-.6-.4-1-1-1s-1 .4-1 1c0 5.5 4.5 10 10 10c2.6 0 5-1 6.9-2.8V21c0 .6.4 1 1 1s1-.4 1-1v-4.5c0-.6-.5-1-1-1z"/>`),
		g.Group(children),
	)
}