package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CentosLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2l4.292 4.292l1.061-1.06L16.121 4H20v3.879l-1.233-1.233l-1.06 1.061L22 12l-4.292 4.293l1.059 1.059L20 16.121V20h-3.88l1.232-1.233l-1.059-1.06L12 22l-4.293-4.293l-1.061 1.06L7.879 20H4v-3.88l1.231 1.232l1.061-1.06L2 12l4.293-4.293l-1.062-1.061L4 7.879V4h3.879L6.646 5.23l1.062 1.062L12 2Zm0 11.413l-2.88 2.879l2.88 2.88l2.879-2.88L12 13.412ZM7.707 9.12L4.828 12l2.878 2.878l2.88-2.88l-2.879-2.877Zm8.585 0l-2.877 2.878l2.878 2.879L19.172 12l-2.88-2.879ZM12 4.828L9.122 7.707l2.879 2.878l2.877-2.879L12 4.828Z"/>`),
		g.Group(children),
	)
}