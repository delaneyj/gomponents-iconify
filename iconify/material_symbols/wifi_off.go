package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.9 11.1q-1.925-1.925-4.462-3.012T12 7q-.525 0-1.012.038T10 7.15L7.45 4.6q1.1-.3 2.238-.45T12 4q3.55 0 6.625 1.325T24 8.95l-2.1 2.15Zm-4 3.95l-.725-.725l-.725-.725l-3.6-3.6q2.025.2 3.788 1.025T19.75 13.2l-1.85 1.85Zm1.85 7.55l-9.4-9.45q-1.175.275-2.187.825T6.35 15.35l-2.1-2.15q.8-.8 1.725-1.4t1.975-1.05L5.7 8.5q-1.025.525-1.913 1.163T2.1 11.1L0 8.95q.8-.8 1.663-1.437T3.5 6.3L1.4 4.2l1.4-1.4l18.4 18.4l-1.45 1.4ZM12 21l-3.525-3.55q.675-.675 1.575-1.063T12 16q1.05 0 1.95.388t1.575 1.062L12 21Z"/>`),
		g.Group(children),
	)
}