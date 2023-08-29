package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gastrointestinal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M23 5c-1.146 4.46-1.146 7.773 0 9.938c1.719 3.247 5.7 5.042.574 9.373c-5.126 4.332-10.323.697-13.532.697c-3.208 0-6.02 3.613-6.02 8.992c0 3.586.66 6.586 1.978 9"/><path stroke-linejoin="round" d="M29.984 5c-1.441 5.329-1.113 8.709.985 10.14c3.148 2.145 3.389-2.336 9.172.33c5.783 2.666 4.28 11.961.953 16.746C37.768 37 30.667 41.835 24.008 41c-6.659-.835-9.135-8.98-9.992-8.98c-.857 0-2.668.025-3.022 3.397c-.236 2.248.435 4.776 2.01 7.583"/><path d="M35.645 28.29c-.578 1.712-1.555 3.048-2.93 4.01c-1.375.962-3.12 1.528-5.237 1.7"/></g>`),
		g.Group(children),
	)
}