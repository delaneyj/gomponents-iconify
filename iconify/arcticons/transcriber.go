package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transcriber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 38.722L33.93 9.278l-9.937 29.444m3.381-9.937h12.882M9.855 14.16a4.128 4.128 0 0 1 8.257 0m0 9.43a4.128 4.128 0 1 1-8.257 0m0-9.43v9.487m8.257-9.487v9.43m-1.008 8.87A9.632 9.632 0 0 1 4.5 24.149m9.536 8.792v5.781"/>`),
		g.Group(children),
	)
}