package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.5 4h-1.6c.1-.4.1-.8.1-1.2c-.1-.3-.2-.6-.4-.9c-.2-.3-.4-.5-.7-.6c-.3-.1-.6-.3-.9-.3c-.3 0-.6 0-.9.2c-.7.2-1.2.7-1.6 1.3c-.4-.6-.9-1.1-1.6-1.3c-.3-.1-.6-.2-.9-.2c-.3 0-.6.1-.9.3c-.3.1-.5.3-.7.6c-.2.2-.3.6-.4.9c0 .4 0 .8.1 1.2H1.5l-.5.5v9l.5.5h12l.5-.5v-9l-.5-.5zM7 13H2V5h5v8zm0-9H4v-.2c-.1-.3-.1-.5-.1-.8c.1-.2.1-.4.3-.5c.1-.2.3-.3.5-.4c.1-.1.3-.1.5-.1s.4 0 .6.1c.3.1.6.3.8.6c.2.3.4.6.4 1V4zm1-.3c0-.4.2-.7.4-1c.2-.3.5-.5.8-.6c.2-.1.4-.1.6-.1c.2 0 .4 0 .6.1c.2.1.3.2.5.4c.1.1.1.3.2.5c0 .3 0 .5-.1.8c0 .1 0 .1-.1.2H8v-.3zm5 9.3H8V5h5v8z"/>`),
		g.Group(children),
	)
}