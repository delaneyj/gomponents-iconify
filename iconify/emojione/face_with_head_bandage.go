package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithHeadBandage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="31.9" cy="32.1" r="29.9" fill="#ffdd67"/><g fill="#664e27"><circle cx="43.7" cy="30" r="5"/><circle cx="20.2" cy="30" r="5"/></g><path fill="#cccfd4" d="M32 2c-6 0-11.7 1.8-16.4 4.9c17.7 9.1 33.2 23 41.7 41.2c3-4.6 4.7-10.2 4.7-16.1C62 15.4 48.6 2 32 2"/><path fill="#fff" d="M32.9 3.6C27.4 3.6 22.3 5.2 18 8c16.7 7.7 31.6 19.3 39.4 35.8c3.1-5.8 2.9-12.9 2.9-12.9C60.2 15.8 48 3.6 32.9 3.6"/><path fill="#cccfd4" d="M15.8 6.7c14.5-.1 28.9 3.3 41.9 9.8C60.4 21 62 26.3 62 32v1.4A92.956 92.956 0 0 0 4.5 20c2.4-5.5 6.4-10.1 11.3-13.3"/><path fill="#fff" d="M17 8.1c13.6-.1 27.6 4 39.8 9.6c3.3 6.5 3.5 11.9 3.5 11.9c-16.5-8.7-35.4-13.1-52.9-11.4c0 0 4.2-7.8 9.6-10.1"/><path fill="#664e27" d="M40.3 42.8c-5.8-1.5-12-.4-16.9 3c-1.2.9 1.1 4 2.3 3.2c3.2-2.3 8.4-3.8 13.7-2.4c1.3.3 2.4-3.3.9-3.8"/>`),
		g.Group(children),
	)
}