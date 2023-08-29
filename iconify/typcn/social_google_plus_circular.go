package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialGooglePlusCircular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.8 13.1l-.4-.3c-.1-.1-.3-.2-.3-.5l.3-.6c.5-.4 1-.8 1-1.7s-.6-1.4-.8-1.6h.7L14 8h-2.4c-.6 0-1.4.1-2.1.6c-.5.4-.8 1-.8 1.6c0 .9.7 1.9 2 1.9h.4c-.1.1-.1.2-.1.4c0 .4.2.6.4.8c-.5 0-1.5.1-2.3.5c-.7.4-.9 1-.9 1.4c0 .9.8 1.7 2.5 1.7c2 0 3.1-1.1 3.1-2.2c0-.7-.5-1.1-1-1.6zm-1.6-1.3c-1 0-1.5-1.3-1.5-2.1c.1-.4.1-.7.3-.9s.5-.4.8-.4c1 0 1.5 1.3 1.5 2.2c0 .2 0 .6-.3.9c-.1.1-.5.3-.8.3zm.1 4.7c-1.3 0-2.1-.6-2.1-1.4c0-.8.8-1.1 1-1.2c.5-.2 1.1-.2 1.2-.2h.3c.9.6 1.3 1 1.3 1.6c0 .7-.6 1.2-1.7 1.2zM15 12h-1v-1h1v-1h1v1h1v1h-1v1h-1zm-3 9c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9zm0-16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7z"/>`),
		g.Group(children),
	)
}