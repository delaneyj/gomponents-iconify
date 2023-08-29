package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightSpokedAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#37b34a" d="M63.738 56.864a6.87 6.87 0 0 1-6.869 6.869H6.87A6.871 6.871 0 0 1 0 56.864V6.869A6.873 6.873 0 0 1 6.87 0h49.999a6.872 6.872 0 0 1 6.869 6.869v49.995z"/><path fill="#0f7b40" d="M.192 48.619v8.282a6.83 6.83 0 0 0 6.825 6.832h49.702a6.828 6.828 0 0 0 6.825-6.832V7.206C52.729 46.506 13.84 48.994.194 48.616"/><path fill="#fff" d="m60.925 32.19l-20.971-5.18L52.44 11.292L36.726 23.774L31.821 2.781L27.01 23.774L11.162 11.468l-.004.004L23.78 27.01L2.814 32.18l20.966 4.539L11.301 52.43L27.01 39.949l4.856 21l4.858-21L52.44 52.43L39.958 36.719L60.921 32.2z"/>`),
		g.Group(children),
	)
}