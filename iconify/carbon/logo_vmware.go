package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoVmware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.4 11C3 10.1 2 9.7 1 10.2s-1.3 1.4-.9 2.3l4 8.5c.6 1.3 1.3 2 2.5 2c1.3 0 1.9-.8 2.5-2c0 0 3.4-7.4 3.4-7.5c0-.1.1-.3.5-.3c.3 0 .5.2.5.6V21c0 1.1.6 2 1.8 2s1.8-.9 1.8-2v-6c0-1.1.8-1.9 1.9-1.9c1.1 0 1.9.8 1.9 1.9v6c0 1.1.6 2 1.8 2s1.8-.9 1.8-2v-6c0-1.1.8-1.9 1.9-1.9c1.1 0 1.9.8 1.9 1.9v6c0 1.1.6 2 1.8 2s1.8-.9 1.8-2v-6.8c0-2.5-2-4.2-4.4-4.2s-3.9 1.7-3.9 1.7c-.8-1-1.9-1.7-3.8-1.7c-2 0-3.7 1.7-3.7 1.7c-.8-1-2.2-1.7-3.3-1.7c-1.7 0-3.1.8-4 2.7l-2.5 5.9L3.4 11"/>`),
		g.Group(children),
	)
}