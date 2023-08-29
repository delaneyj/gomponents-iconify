package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashicorp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m305.778 379.733l76.8-42.666V45.511L305.778 0v217.6h-98.134v-85.333l-78.222 42.666v291.556L207.644 512V294.4h98.134z"/><path fill="currentColor" d="M419.556 65.422v295.822l-113.778 62.578V512l184.889-106.667V106.667zM207.644 0L21.333 106.667v298.666l72.534 41.245V150.756l113.777-62.578z"/>`),
		g.Group(children),
	)
}