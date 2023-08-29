package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deepl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.496 11.972L25.034 4.777a2.068 2.068 0 0 0-2.068 0l-12.462 7.195a2.068 2.068 0 0 0-1.034 1.79v14.54a2.068 2.068 0 0 0 .84 1.664l18.074 13.33a1.034 1.034 0 0 0 1.648-.832V34.85a1.034 1.034 0 0 1 .517-.896l6.946-4.01a2.068 2.068 0 0 0 1.035-1.792v-14.39a2.068 2.068 0 0 0-1.034-1.79Z"/><circle cx="31.375" cy="21.538" r="3.093" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.002" cy="15.352" r="3.093" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.002" cy="27.725" r="3.093" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.766 26.342l4.295-2.147m-4.295-7.461l6.844 3.422"/>`),
		g.Group(children),
	)
}