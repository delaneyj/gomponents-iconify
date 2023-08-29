package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Erlang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M87.2 53.5H0v405h100.4c-49.7-52.6-78.8-125.3-78.7-212.1c-.1-76.7 24-142.7 65.5-192.9zm238.2 9.7c-45.9.1-85.1 33.5-89.2 83.2h169.9c-1.1-49.7-34.5-83.1-80.7-83.2zm230.7-9.6h.3l-.1-.1zm.3 0c31.4 42.7 48.7 97.5 46.2 162.7c.5 6 .5 11.7 0 24.1H230.2c-.2 109.7 38.9 194.9 138.6 195.3c68.5-.3 118-51 151.9-106.1l96.4 48.2c-17.4 30.9-36.5 57.8-57.9 80.8H640v-405z"/>`),
		g.Group(children),
	)
}