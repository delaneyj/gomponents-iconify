package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Up(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.5 6a.47.47 0 0 1-.35-.15L8 1.71L3.85 5.85c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71L7.65.65c.2-.2.51-.2.71 0l4.5 4.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		g.Group(children),
	)
}