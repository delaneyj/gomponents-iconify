package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twrp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34 5.59v4.9h5.77a.48.48 0 0 1 .48.49a.52.52 0 0 1-.09.29L28.77 26.49a.49.49 0 0 1-.69.09a.43.43 0 0 1-.1-.09L16.63 11.28a.48.48 0 0 1 .37-.78h5.77V2.93a.4.4 0 0 0-.41-.39h0a21.53 21.53 0 0 0-8.81 40.27a.41.41 0 0 0 .59-.35v-5.11h-5.8a.49.49 0 0 1-.49-.49a.54.54 0 0 1 .1-.3l11.36-15.24a.5.5 0 0 1 .7-.08l.09.08l11.35 15.24a.49.49 0 0 1-.1.69a.53.53 0 0 1-.29.1h-5.77v7.71a.4.4 0 0 0 .4.4h0a21.53 21.53 0 0 0 8.84-40.22a.41.41 0 0 0-.53.21a.39.39 0 0 0 0 .14Z"/>`),
		g.Group(children),
	)
}