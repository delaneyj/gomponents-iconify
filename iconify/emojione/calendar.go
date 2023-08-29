package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ba9372" d="M58.1 19.6c-9.8-6-25-14.8-27.1-14.8c-2.1 0-16.7 8.9-26 14.8l-.5-.9C6.9 17.2 28 3.8 31 3.8s25.1 13.4 27.6 15l-.5.8"/><path fill="#93a2aa" d="M62 56.8c0 1.6-1.2 3-2.7 3H6.8c-1.5 0-2.7-1.3-2.7-3V21.1c0-1.6 1.2-3 2.7-3h52.5c1.5 0 2.7 1.3 2.7 3v35.7"/><path fill="#ed4c5c" d="M60 21.1c0-1.6-1.2-3-2.7-3H4.7c-1.5 0-2.7 1.3-2.7 3v9.5h58v-9.5z"/><path fill="#d9e3e8" d="M2 30.6v26.2c0 1.6 1.2 3 2.7 3h52.5c1.5 0 2.7-1.3 2.7-3V30.6H2z"/><path fill="#93a2aa" d="M4.5 33h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zM28 37.4h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm-47 4.4h6V44h-6zm7.8 0h6V44h-6zm7.9 0h6V44h-6zm7.8 0h6V44h-6zm7.8 0h6V44h-6zm7.9 0h6V44h-6zm7.8 0h6V44h-6zm-47 4.5h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm-47 4.4h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm-47 4.4h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.8 0h6v2.2h-6zm7.9 0h6v2.2h-6z"/><ellipse cx="31.2" cy="6.2" fill="#333" rx="1.8" ry="1.9"/><ellipse cx="31" cy="6.2" fill="#93a2aa" rx="1.8" ry="1.9"/><path fill="#fff" d="M19.5 25.5v.2c0 .6.1 1 .2 1.3c.1.2.3.4.7.4c.4 0 .6-.1.7-.4c.1-.2.1-.4.1-.8v-5.5h1.6v5.4c0 .7-.1 1.2-.3 1.6c-.4.7-1 1-2 1s-1.6-.3-2-.8c-.3-.5-.5-1.3-.5-2.2v-.2h1.5m5-4.8h1.6v4.8c0 .5.1.9.2 1.2c.2.4.6.7 1.3.7c.6 0 1.1-.2 1.3-.7c.1-.3.1-.7.1-1.2v-4.8h1.6v4.8c0 .8-.1 1.5-.4 1.9c-.5.8-1.4 1.3-2.7 1.3c-1.3 0-2.2-.4-2.7-1.3c-.3-.5-.4-1.1-.4-1.9c.1 0 .1-4.8.1-4.8m7.7 0h1.6v6.4h3.8v1.4h-5.4v-7.8m10 0H44l-2.6 4.9v2.9h-1.6v-2.9l-2.7-4.9H39l1.6 3.4l1.6-3.4"/>`),
		g.Group(children),
	)
}