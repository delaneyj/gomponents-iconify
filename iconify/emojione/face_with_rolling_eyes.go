package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithRollingEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><path fill="#fff" d="M45.5 40c6.1 0 11-4.9 11-11c0-1.6-.4-3.2-1-4.6c-2.8-1.3-6.3-2-10-2s-7.2.8-10 2c-.6 1.4-1 2.9-1 4.6c0 6.1 4.9 11 11 11"/><path fill="#664e27" d="M46 22.4c-.5.7-.8 1.6-.8 2.5c0 2.5 2 4.5 4.5 4.5s4.5-2 4.5-4.5c0-.4-.1-.7-.1-1c-2.4-.9-5.2-1.5-8.1-1.5"/><path fill="#fff" d="M18.5 40c6.1 0 11-4.9 11-11c0-1.6-.4-3.2-1-4.6c-2.8-1.3-6.3-2-10-2s-7.2.8-10 2c-.6 1.4-1 2.9-1 4.6c0 6.1 4.9 11 11 11"/><path fill="#664e27" d="M19 22.4c-.5.7-.8 1.6-.8 2.5c0 2.5 2 4.5 4.5 4.5s4.5-2 4.5-4.5c0-.4-.1-.7-.1-1c-2.4-.9-5.2-1.5-8.1-1.5m21.6 24.9c-5.4-2.5-11.8-2.5-17.2 0c-1.4.7.3 4.2 1.6 3.5c3.6-1.7 8.9-2.3 13.9 0c1.4.6 3.1-2.8 1.7-3.5"/>`),
		g.Group(children),
	)
}