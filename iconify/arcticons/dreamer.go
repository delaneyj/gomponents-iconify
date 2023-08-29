package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dreamer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.769 39.615a7.447 7.447 0 0 1-6.166 3.885a6.059 6.059 0 0 1-6.167-3.885c-.813-1.734-2.372-5.233-2.372-5.233l-8.883-.005a6.027 6.027 0 1 1 0-12.054s.99-6.13 6.137-6.13h4.02l8.656-6.958V20.21l4.774 12.292c1.026 2.447 1.273 4.912 0 7.113Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.476 11.848h-7.575c-3.974 0-5.938 2.697-5.748 6.24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.556 11.903L13.666 4.5l3.706 9.13m15.693 1.803s4.394 7.258 6.218 12.61c1.591 4.667-.66 5.583-.66 5.583"/><circle cx="20.301" cy="22.575" r="1.474" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}