package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spritpreise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 28.856l10.447-10.48L25.288 28.75L14.861 39.21L4.5 28.856ZM40.812 14.8a2.687 2.687 0 0 1 2.683 2.691a2.688 2.688 0 0 1-2.683 2.692a2.688 2.688 0 0 1-2.684-2.692a2.688 2.688 0 0 1 2.684-2.692h0ZM26.87 12.053v3.834h4.009l-4.897 7.628v-4.243l-3.75-.091l4.637-7.127Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.844 15.733l-.352-.371l-5.137-5.42"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.359 31.896v-3.69h7.07"/><path d="m15.913 30.722l2.516-2.516l-2.516-2.515"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.58 22.02V11.244a2.454 2.454 0 0 1 2.453-2.454h11.044a2.454 2.454 0 0 1 2.454 2.454v24.541a2.454 2.454 0 0 1-2.454 2.454H21.033a2.454 2.454 0 0 1-2.454-2.454v-.305"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 17.491v15.274a2.686 2.686 0 0 1-2.686 2.686h0a2.686 2.686 0 0 1-2.686-2.686v-7.532c0-.95-.769-1.718-1.718-1.718h-1.88"/>`),
		g.Group(children),
	)
}