package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.16 8.489c-2.709-3.165-7.619-3.169-10.332-.007a75.064 75.064 0 0 0-8.087 11.427A76.399 76.399 0 0 0 4.89 32.652c-1.38 3.917 1.063 8.155 5.144 8.923a75.073 75.073 0 0 0 13.965 1.295a76.408 76.408 0 0 0 13.967-1.304c4.079-.764 6.526-4.995 5.152-8.911a75.07 75.07 0 0 0-5.863-12.746A76.403 76.403 0 0 0 29.16 8.49Z"/><ellipse cx="20.663" cy="21.139" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.337" ry="3.971"/><ellipse cx="27.337" cy="21.139" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.337" ry="3.971"/><circle cx="20.663" cy="21.139" r=".75" fill="currentColor"/><circle cx="27.267" cy="21.139" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.554 24.367c-1.915 3.38-.789 7.706 3.492 5.453"/>`),
		g.Group(children),
	)
}