package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QbittorrentRemote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.651 22.364a5.034 5.034 0 0 1 5.035-5.035h0a5.034 5.034 0 0 1 5.034 5.035v3.272a5.034 5.034 0 0 1-5.034 5.035h0a5.034 5.034 0 0 1-5.035-5.035m0 5.035V10.533m-5.302 15.103a5.034 5.034 0 0 1-5.035 5.035h0a5.034 5.034 0 0 1-5.034-5.035v-3.272a5.034 5.034 0 0 1 5.034-5.035h0a5.034 5.034 0 0 1 5.035 5.035m0-5.035v20.138"/>`),
		g.Group(children),
	)
}