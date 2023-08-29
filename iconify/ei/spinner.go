package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 18c-.6 0-1-.4-1-1V9c0-.6.4-1 1-1s1 .4 1 1v8c0 .6-.4 1-1 1z"/><path fill="currentColor" d="M25 42c-.6 0-1-.4-1-1v-8c0-.6.4-1 1-1s1 .4 1 1v8c0 .6-.4 1-1 1zm4-23c-.2 0-.3 0-.5-.1c-.4-.3-.6-.8-.3-1.3l4-6.9c.3-.4.8-.6 1.3-.3c.4.3.6.8.3 1.3l-4 6.9c-.2.2-.5.4-.8.4zM17 39.8c-.2 0-.3 0-.5-.1c-.4-.3-.6-.8-.3-1.3l4-6.9c.3-.4.8-.6 1.3-.3c.4.3.6.8.3 1.3l-4 6.9c-.2.2-.5.4-.8.4z" opacity=".3"/><path fill="currentColor" d="M21 19c-.3 0-.6-.2-.8-.5l-4-6.9c-.3-.4-.1-1 .3-1.3c.4-.3 1-.1 1.3.3l4 6.9c.3.4.1 1-.3 1.3c-.2.2-.3.2-.5.2z" opacity=".93"/><path fill="currentColor" d="M33 39.8c-.3 0-.6-.2-.8-.5l-4-6.9c-.3-.4-.1-1 .3-1.3c.4-.3 1-.1 1.3.3l4 6.9c.3.4.1 1-.3 1.3c-.2.1-.3.2-.5.2z" opacity=".3"/><path fill="currentColor" d="M17 26H9c-.6 0-1-.4-1-1s.4-1 1-1h8c.6 0 1 .4 1 1s-.4 1-1 1z" opacity=".65"/><path fill="currentColor" d="M41 26h-8c-.6 0-1-.4-1-1s.4-1 1-1h8c.6 0 1 .4 1 1s-.4 1-1 1z" opacity=".3"/><path fill="currentColor" d="M18.1 21.9c-.2 0-.3 0-.5-.1l-6.9-4c-.4-.3-.6-.8-.3-1.3c.3-.4.8-.6 1.3-.3l6.9 4c.4.3.6.8.3 1.3c-.2.3-.5.4-.8.4z" opacity=".86"/><path fill="currentColor" d="M38.9 33.9c-.2 0-.3 0-.5-.1l-6.9-4c-.4-.3-.6-.8-.3-1.3c.3-.4.8-.6 1.3-.3l6.9 4c.4.3.6.8.3 1.3c-.2.3-.5.4-.8.4z" opacity=".3"/><path fill="currentColor" d="M11.1 33.9c-.3 0-.6-.2-.8-.5c-.3-.4-.1-1 .3-1.3l6.9-4c.4-.3 1-.1 1.3.3c.3.4.1 1-.3 1.3l-6.9 4c-.1.2-.3.2-.5.2z" opacity=".44"/><path fill="currentColor" d="M31.9 21.9c-.3 0-.6-.2-.8-.5c-.3-.4-.1-1 .3-1.3l6.9-4c.4-.3 1-.1 1.3.3c.3.4.1 1-.3 1.3l-6.9 4c-.2.2-.3.2-.5.2z" opacity=".3"/>`),
		g.Group(children),
	)
}