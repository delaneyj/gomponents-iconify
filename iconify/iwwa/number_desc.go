package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberDesc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M24.195 32.164h-2.644a.75.75 0 0 0 0 1.5h2.644a.75.75 0 0 0 0-1.5zm2.909-6.414h-5.552a.75.75 0 0 0 0 1.5h5.552a.75.75 0 0 0 0-1.5zm7.908-12.828h-13.46a.75.75 0 0 0 0 1.5h13.46a.75.75 0 0 0 0-1.5zm3.321-6.414H21.552a.75.75 0 0 0 0 1.5h16.781a.75.75 0 0 0 0-1.5zM17.468 24.455l-6.692 8.076V6.934a.75.75 0 0 0-1.5 0v25.563l-6.693-8.043a.75.75 0 1 0-1.154.959l8.037 9.657a.75.75 0 0 0 .577.271h.001c.223 0 .434-.1.577-.271l8.003-9.657a.75.75 0 0 0-1.156-.958zm13.544-5.119h-9.46a.75.75 0 0 0 0 1.5h9.46a.75.75 0 0 0 0-1.5z"/>`),
		g.Group(children),
	)
}