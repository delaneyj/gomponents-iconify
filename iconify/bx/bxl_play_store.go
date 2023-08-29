package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxlPlayStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12.954 11.616l2.957-2.957L6.36 3.291c-.633-.342-1.226-.39-1.746-.016l8.34 8.341zm3.461 3.462l3.074-1.729c.6-.336.929-.812.929-1.34c0-.527-.329-1.004-.928-1.34l-2.783-1.563l-3.133 3.132l2.841 2.84zM4.1 4.002c-.064.197-.1.417-.1.658v14.705c0 .381.084.709.236.97l8.097-8.098L4.1 4.002zm8.854 8.855L4.902 20.91c.154.059.32.09.495.09c.312 0 .637-.092.968-.276l9.255-5.197l-2.666-2.67z" fill="currentColor"/>`),
		g.Group(children),
	)
}