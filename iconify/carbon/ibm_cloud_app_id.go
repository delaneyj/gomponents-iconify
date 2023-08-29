package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudAppId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 18h-6a3 3 0 0 0-3 3v2h2v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h2v-2a3 3 0 0 0-3-3zm-3-1a4 4 0 1 0 0-8a4 4 0 0 0 0 8zm0-6a2 2 0 1 1 0 4a2 2 0 0 1 0-4z"/><path fill="currentColor" d="M17 30C9.28 30 3 23.72 3 16h2c0 6.617 5.383 12 12 12c5.226 0 9.816-3.338 11.421-8.307l1.904.614A13.961 13.961 0 0 1 17 30zm14-14h-2c0-6.617-5.383-12-12-12V2c7.72 0 14 6.28 14 14zM6 14l-2.139-1.013A5.022 5.022 0 0 1 1 8.467V2h10v6.468a5.021 5.021 0 0 1-2.861 4.52L6 14zM3 4v4.468a3.01 3.01 0 0 0 1.717 2.71L6 11.787l1.283-.607A3.012 3.012 0 0 0 9 8.468V4H3z"/>`),
		g.Group(children),
	)
}