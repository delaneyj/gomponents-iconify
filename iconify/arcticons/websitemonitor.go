package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Websitemonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.226 25.999A16.093 16.093 0 0 0 39.37 24a16.091 16.091 0 0 0-.143-1.999l4.335-3.393a1.022 1.022 0 0 0 .246-1.312l-4.1-7.103a1.014 1.014 0 0 0-1.25-.44l-5.105 2.06a15.554 15.554 0 0 0-3.464-2.02l-.769-5.432a1.054 1.054 0 0 0-1.025-.861h-8.2a1.035 1.035 0 0 0-1.015.861l-.768 5.433a15.167 15.167 0 0 0-3.465 2.019l-5.104-2.06a1.025 1.025 0 0 0-1.251.44l-4.1 7.104a1.027 1.027 0 0 0 .246 1.312l4.326 3.393A16.08 16.08 0 0 0 8.62 24a16.093 16.093 0 0 0 .143 1.999l-4.325 3.393a1.022 1.022 0 0 0-.246 1.312l4.1 7.103a1.014 1.014 0 0 0 1.25.44l5.105-2.06a15.554 15.554 0 0 0 3.464 2.02l.769 5.432a1.035 1.035 0 0 0 1.015.861h8.2a1.035 1.035 0 0 0 1.015-.861l.768-5.433a15.17 15.17 0 0 0 3.465-2.019l5.104 2.06a1.025 1.025 0 0 0 1.25-.44l4.1-7.104a1.027 1.027 0 0 0-.246-1.312Z"/><circle cx="24" cy="24" r="5.625" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.375A12.938 12.938 0 0 0 13.5 24A12.938 12.938 0 0 0 24 29.625A12.938 12.938 0 0 0 34.5 24A12.938 12.938 0 0 0 24 18.375Z"/><circle cx="24" cy="24" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}