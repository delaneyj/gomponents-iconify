package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithThermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><g fill="#664e27"><circle cx="43.5" cy="30.3" r="5"/><circle cx="20.5" cy="30.3" r="5"/></g><path fill="#917524" d="M25.6 15.2c-3.2 2.7-7.5 3.9-11.7 3.1c-.6-.1-1.1 2-.4 2.2c4.8.9 9.8-.5 13.5-3.6c.5-.5-1-2.1-1.4-1.7m24.5 3c-4.2.7-8.5-.4-11.7-3.1c-.4-.4-2 1.2-1.4 1.7c3.7 3.2 8.7 4.5 13.5 3.6c.7-.2.2-2.3-.4-2.2"/><path fill="#e5e5e5" d="m4.4 38.7l24.2 14.2l4.6-6.5l-21.4-18.1c-6.4-4.5-13.7 5.9-7.4 10.4"/><path fill="#fff" d="m5.8 36.5l23.8 15l2.8-3.9l-22.2-17.3c-3.7-2.6-8.1 3.6-4.4 6.2"/><path fill="#ed4c5c" d="m7.5 36.2l22.1 14.4l1.9-2.6l-21.1-16z"/><ellipse cx="9" cy="34.1" fill="#a5203c" rx="2.6" ry=".9" transform="rotate(-54.808 8.95 34.089)"/><path fill="#51575b" d="m18.1 41.5l-2 2.8l-.9-.5l2-2.9zm2.4 1.7l-1.9 2.7l-.9-.5l1.9-2.8zm2.5 1.7l-1.9 2.6l-.8-.5l1.8-2.6zm2.4 1.7L23.8 49l-.9-.6l1.7-2.3zm2.5 1.7l-1.5 2.2l-.9-.6l1.5-2.1z"/><path fill="#664e27" d="M41.5 43.9c-6 .3-11.6 3.3-15.2 8c-.9 1.2 2.2 3.5 3.1 2.3c2.4-3.1 6.8-6.2 12.3-6.5c1.4 0 1.3-3.8-.2-3.8"/>`),
		g.Group(children),
	)
}