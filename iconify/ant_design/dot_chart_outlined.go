package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotChartOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M888 792H200V168c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h752c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM288 604a64 64 0 1 0 128 0a64 64 0 1 0-128 0zm118-224a48 48 0 1 0 96 0a48 48 0 1 0-96 0zm158 228a96 96 0 1 0 192 0a96 96 0 1 0-192 0zm148-314a56 56 0 1 0 112 0a56 56 0 1 0-112 0z"/>`),
		g.Group(children),
	)
}