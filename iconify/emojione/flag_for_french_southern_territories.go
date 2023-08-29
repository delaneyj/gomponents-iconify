package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFrenchSouthernTerritories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M39.1 2.9v24.7H2.3C2.1 29 2 30.5 2 32c0 16.6 13.4 30 30 30s30-13.4 30-30c0-14.1-9.8-25.9-22.9-29.1z"/><path fill="#fff" d="M39.1 2.9c-.2-.1-.4-.1-.6-.1V27H25.6V2.7C20.8 3.7 16.4 6 12.7 9v18H2.4c0 .2-.1.4-.1.6h36.8V2.9"/><path fill="#2a5f9e" d="M2.4 27h10.3V9C7.4 13.5 3.6 19.8 2.4 27z"/><path fill="#ed4c5c" d="M25.6 2.7V27h12.9V2.7C36.4 2.2 34.2 2 32 2s-4.3.2-6.4.7"/><g fill="#fff"><path d="M42.8 43.8v-4.7h1.9l1.4-2.3h-3.4v-1.5H47l1.5-2.4H34.3l1.5 2.4h4.3v8.4l-3.2-5.1l-4.4 6.9h2.1l.4-.8h3.6l2.8 5l2.8-5h3.6l.4.8h2.1l-4.4-6.9l-3.1 5.2m-6.8-.6l.9-1.6l.9 1.6H36m9.2 0l.9-1.6l.9 1.6h-1.8m-3.2 10l-.5-1.7l-.6 1.7h-1.8l1.5 1.1L40 56l1.5-1.1l1.4 1.1l-.5-1.7l1.4-1.1zm-5.3-4l-.5-1.7l-.6 1.7h-1.8l1.5 1.1l-.6 1.7l1.5-1.1l1.4 1.1l-.5-1.7l1.4-1.1zm10.6 0l-.5-1.7l-.6 1.7h-1.8l1.5 1.1l-.6 1.7l1.5-1.1l1.4 1.1l-.5-1.7l1.4-1.1z"/><path d="m31.7 39l1.5 1.1l-.6-1.7l1.5-1.1h-1.8l-.6-1.7l-.5 1.7h-1.8l1.4 1.1l-.5 1.7zm20.5-.6l1.4-1.1h-1.8l-.5-1.7l-.6 1.7h-1.8l1.5 1.1l-.6 1.7l1.5-1.1l1.4 1.1z"/></g>`),
		g.Group(children),
	)
}