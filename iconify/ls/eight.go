package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M343 349c67 40 112 113 112 196c0 126-102 227-227 227C102 772 0 671 0 545c0-83 44-156 111-196c-47-35-78-91-78-154C33 87 120 0 228 0s193 87 193 195c0 63-30 119-78 154zM105 195c0 68 55 123 123 123s122-55 122-123S296 72 228 72s-123 55-123 123zm123 507c86 0 155-71 155-157s-69-156-155-156S71 459 71 545s71 157 157 157z"/>`),
		g.Group(children),
	)
}