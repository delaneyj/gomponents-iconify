package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 30.088v.596a6.287 6.287 0 0 1-6.301 6.302H13.44c-.34 0-.681 0-1.022-.085h-.426c-4.172 0-7.493-3.407-7.493-7.579c0-3.917 2.98-7.153 6.812-7.494c1.192-6.13 6.642-10.814 13.199-10.814c7.238 0 13.113 5.705 13.454 12.773c3.066.34 5.535 3.065 5.535 6.301Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.28 30.77a7.016 7.016 0 0 0 1.873-3.492c.681-3.661-1.703-7.323-5.365-8.004m-6.216 1.873A7.016 7.016 0 0 0 17.7 24.64c-.682 3.661 1.703 7.238 5.364 8.004"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.875 31.366l1.789 1.788h0h-4.939v-4.939l1.788 1.788l1.362 1.363zm-10.899-10.9l-1.788-1.788h4.939v5.024l-1.789-1.873l-1.362-1.363z"/>`),
		g.Group(children),
	)
}