package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M32 14a8 8 0 1 1-16 0a8 8 0 0 1 16 0ZM16.212 31.011a2 2 0 0 0-2.117 2.601l2.117-2.6Zm8.583 4.236A3.501 3.501 0 0 1 24 41.965V42h-8a9 9 0 1 1 0-18h16a9 9 0 1 1 0 18h-4.257a5.502 5.502 0 0 0-2.175-8.598l-.773 1.845ZM23.035 26h3.463l-5.86 14H16a6.968 6.968 0 0 1-3.677-1.042l9.647-11.852L23.036 26ZM34 33a2 2 0 0 1-2 2v-4a2 2 0 0 1 2 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}