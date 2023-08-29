package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EslintOld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#3A33D1" d="M100.034 262.106L.598 172.57L28.45 41.694L155.626.354l99.436 89.535l-27.851 130.876l-127.177 41.34Zm-48.086-106.18l59.291 53.307l75.828-24.695l16.645-78.004l-59.291-53.417l-75.828 24.805l-16.645 78.003Z"/><path fill="#6464E2" d="M181.301 223.92H74.359l-53.525-92.69l53.525-92.69h106.942l53.525 92.69l-53.525 92.69ZM93.18 191.283h69.3l34.705-60.053l-34.705-60.053h-69.3L58.584 131.23l34.596 60.053Z"/>`),
		g.Group(children),
	)
}