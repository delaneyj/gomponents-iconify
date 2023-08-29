package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAnsibleCommunity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.8 27.7L17.5 3.1c-.3-.7-.9-1.1-1.6-1.1c-.7 0-1.3.4-1.6 1.1L3 30h3.9l4.5-11.1l13.3 10.7c.5.4.9.6 1.4.6c1 0 1.9-.7 1.9-1.8c0-.2-.1-.4-.2-.7zM15.9 7.5l6.7 16.4L12.5 16l3.4-8.5z"/>`),
		g.Group(children),
	)
}