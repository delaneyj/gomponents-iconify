package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 1a.5.5 0 0 0 0 1c.922 0 1.54.23 1.92.564c.373.325.58.802.58 1.436v3H5.75a.5.5 0 0 0 0 1H7v3c0 .634-.207 1.11-.58 1.436c-.38.334-.998.564-1.92.564a.5.5 0 0 0 0 1c1.078 0 1.96-.27 2.58-.811c.162-.142.302-.3.42-.47c.118.17.258.328.42.47c.62.541 1.502.811 2.58.811a.5.5 0 0 0 0-1c-.922 0-1.54-.23-1.92-.564C8.206 12.111 8 11.634 8 11V8h1.25a.5.5 0 0 0 0-1H8V4c0-.634.207-1.11.58-1.436C8.96 2.23 9.577 2 10.5 2a.5.5 0 0 0 0-1c-1.078 0-1.96.27-2.58.811c-.162.142-.302.3-.42.47a2.588 2.588 0 0 0-.42-.47C6.46 1.27 5.577 1 4.5 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}