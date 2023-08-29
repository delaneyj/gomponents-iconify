package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocketchat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 24.1a10.48 10.48 0 0 0-1.7-5.6a13 13 0 0 0-4.3-4.1a24.17 24.17 0 0 0-13.1-3.5a25.22 25.22 0 0 0-4.8.4a19.92 19.92 0 0 0-3.3-2.4a13.42 13.42 0 0 0-11.8-.1s5 4.1 4.2 7.6A10.84 10.84 0 0 0 5.3 24h0a10.67 10.67 0 0 0 3.4 7.6c.8 3.6-4.2 7.6-4.2 7.6a13.23 13.23 0 0 0 11.7-.1a15.47 15.47 0 0 0 3.3-2.4a25.92 25.92 0 0 0 4.8.4a24.17 24.17 0 0 0 13.1-3.5a15.41 15.41 0 0 0 4.3-4.1a8.57 8.57 0 0 0 1.8-5.4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.69 16.39C12.17 13 18 10.83 24.76 10.9"/><circle cx="16.7" cy="24.5" r="2.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.3" cy="24.5" r="2.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.9" cy="24.5" r="2.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}