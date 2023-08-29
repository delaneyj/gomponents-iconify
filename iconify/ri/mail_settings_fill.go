package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSettingsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13.341A6 6 0 0 0 14.341 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v9.341Zm-9.94-1.658L5.648 6.238L4.353 7.762l7.72 6.555l7.581-6.56l-1.308-1.513l-6.285 5.439Zm4.99 7.866a3.023 3.023 0 0 1 0-1.098l-1.014-.585l1-1.732l1.014.586c.278-.238.599-.425.95-.55V15h2v1.17c.351.125.672.312.95.55l1.014-.586l1 1.732l-1.014.585a3.023 3.023 0 0 1 0 1.098l1.014.585l-1 1.732l-1.014-.586a2.997 2.997 0 0 1-.95.55V23h-2v-1.17a2.997 2.997 0 0 1-.95-.55l-1.014.586l-1-1.732l1.014-.585ZM20 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}