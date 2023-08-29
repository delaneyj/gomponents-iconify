package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedMailboxWithRaisedFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.42.37c-.67 0-1.21.54-1.21 1.21v8.904A2.419 2.419 0 0 0 18.42 15a2.42 2.42 0 0 0 1.21-4.516V5h4.502c.475 0 .868-.369.868-.816V1.816c0-.447-.393-.816-.869-.816h-4.648A1.208 1.208 0 0 0 18.42.37ZM9.27 20H4.98c-.54 0-.98.44-.98.98s.44.98.98.98h4.29c.54 0 .98-.44.98-.98a.969.969 0 0 0-.98-.98Z"/><path d="M7.84 6h8.37v2h-3.542a6.823 6.823 0 0 1 2.002 4.84v9.77h11.32a.5.5 0 0 0 .5-.5v-9.269A4.847 4.847 0 0 0 21.65 8h-1.02V6h1.02c3.774 0 6.828 3.059 6.84 6.837v9.273a2.5 2.5 0 0 1-2.5 2.5h-9.18v5.43h-4.98v-5.43H3.5a2.5 2.5 0 0 1-2.5-2.5v-9.27A6.834 6.834 0 0 1 7.84 6Zm0 2A4.834 4.834 0 0 0 3 12.84v9.27a.5.5 0 0 0 .5.5h9.17v-9.77A4.832 4.832 0 0 0 7.84 8Z"/></g>`),
		g.Group(children),
	)
}