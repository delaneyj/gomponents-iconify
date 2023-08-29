package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClubSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M47 18.875c-.044 0-.086.008-.13.008c.076-.619.13-1.244.13-1.883c0-8.283-6.716-15-15-15c-8.286 0-15 6.717-15 15c0 .639.053 1.264.13 1.883c-.044 0-.086-.008-.13-.008c-8.286 0-15 6.717-15 15s6.714 15 15 15c4.807 0 9.075-2.271 11.819-5.788l-.569 4.851c-.505 3.589-2.994 6.872-6.563 7.5L17 56.375V62h30v-5.625l-4.688-.938c-3.57-.628-6.06-3.909-6.563-7.5l-.571-4.851c2.744 3.518 7.013 5.788 11.821 5.788c8.284 0 15-6.717 15-15S55.284 18.875 47 18.875"/>`),
		g.Group(children),
	)
}