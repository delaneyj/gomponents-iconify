package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagFiveLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M3.794 14.97c.537 2.687.806 4.03 1.693 4.895c.164.16.342.305.53.434C7.04 21 8.41 21 11.15 21h1.703c2.739 0 4.108 0 5.13-.7c.19-.13.367-.276.53-.435c.888-.865 1.157-2.208 1.694-4.894c.771-3.856 1.157-5.784.269-7.15a4 4 0 0 0-.56-.683C18.75 6 16.785 6 12.853 6H11.15c-3.934 0-5.9 0-7.066 1.138a4 4 0 0 0-.559.683c-.888 1.366-.502 3.294.27 7.15Z"/><circle cx="15" cy="10" r="1" fill="currentColor"/><circle cx="9" cy="10" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 6V5a3 3 0 1 1 6 0v1"/></g>`),
		g.Group(children),
	)
}