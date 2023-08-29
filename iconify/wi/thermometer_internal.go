package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerInternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M12.48 19.56c0 .71.24 1.32.73 1.82s1.07.75 1.76.75s1.28-.25 1.79-.75s.76-1.11.76-1.81c0-.63-.22-1.19-.65-1.67c-.43-.48-.96-.77-1.57-.85V9.69c0-.06-.03-.13-.1-.19a.299.299 0 0 0-.22-.1c-.09 0-.16.03-.21.08c-.05.06-.08.12-.08.21v7.34c-.61.09-1.13.37-1.56.85c-.44.49-.65 1.04-.65 1.68z"/>`),
		g.Group(children),
	)
}