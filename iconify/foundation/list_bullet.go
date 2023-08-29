package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M88.721 20.13H26.258a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407zm0 24.892H26.258a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407zm0 24.891H26.258a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407V73.32a3.408 3.408 0 0 0-3.407-3.407z"/><circle cx="12.856" cy="25.108" r="4.984" fill="currentColor"/><circle cx="12.856" cy="49.002" r="4.984" fill="currentColor"/><circle cx="12.856" cy="74.891" r="4.984" fill="currentColor"/>`),
		g.Group(children),
	)
}