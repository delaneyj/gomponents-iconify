package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.573" cy="14.408" r="9.309" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.573" cy="14.251" r="4.515" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.344 23.789l-.072 14.045c-4.271-.972-7.874-2.772-9.793-6.66L14.17 28.5l-9.67.018l.071 9.511l2.87-2.78C11.12 40.62 16.78 42.957 24 42.899m2.763-19.606l-.036 14.54c4.272-.971 7.875-2.771 9.794-6.659L33.83 28.5l9.67.018l-.071 9.511l-2.87-2.78C36.88 40.62 31.22 42.957 24 42.899"/>`),
		g.Group(children),
	)
}