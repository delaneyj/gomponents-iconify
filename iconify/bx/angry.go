package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/><path fill="currentColor" d="M12 14c-3 0-4 3-4 3h8s-1-3-4-3zm-2.439-2.439c.014-.014.023-.03.037-.044l1.031.413l.742-1.857l-5-2l-.742 1.856l1.373.549L7 10.5a1.499 1.499 0 0 0 2.561 1.061zm3.068-1.49l.742 1.857l1.037-.415c.011.011.019.024.029.035a1.488 1.488 0 0 0 2.112 0c.271-.271.438-.644.438-1.056l-.001-.01l1.386-.554l-.742-1.857l-5.001 2z"/>`),
		g.Group(children),
	)
}