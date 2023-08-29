package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#A6704C" d="M16.331 2c-2.57 0-4.98 1.28-6.42 3.4l-.13.19A21.923 21.923 0 0 0 6.101 20l.08.8c.5 5.21 4.88 9.18 10.11 9.18c5.25 0 9.63-3.99 10.11-9.22l.09-.93a21.9 21.9 0 0 0-3.78-14.48A7.772 7.772 0 0 0 16.331 2Z"/><path fill="url(#f434id0)" d="M16.331 2c-2.57 0-4.98 1.28-6.42 3.4l-.13.19A21.923 21.923 0 0 0 6.101 20l.08.8c.5 5.21 4.88 9.18 10.11 9.18c5.25 0 9.63-3.99 10.11-9.22l.09-.93a21.9 21.9 0 0 0-3.78-14.48A7.772 7.772 0 0 0 16.331 2Z"/><path fill="url(#f434id1)" d="M16.331 2c-2.57 0-4.98 1.28-6.42 3.4l-.13.19A21.923 21.923 0 0 0 6.101 20l.08.8c.5 5.21 4.88 9.18 10.11 9.18c5.25 0 9.63-3.99 10.11-9.22l.09-.93a21.9 21.9 0 0 0-3.78-14.48A7.772 7.772 0 0 0 16.331 2Z"/><defs><radialGradient id="f434id0" cx="0" cy="0" r="1" gradientTransform="rotate(89.708 6.161 16.62) scale(24.5003 16.2938)" gradientUnits="userSpaceOnUse"><stop offset=".092" stop-color="#FFCA7F"/><stop offset=".562" stop-color="#E5945B"/><stop offset=".838" stop-color="#B87241"/><stop offset="1" stop-color="#9A6240"/></radialGradient><radialGradient id="f434id1" cx="0" cy="0" r="1" gradientTransform="rotate(97.431 3.208 12.778) scale(22.5645 27.4892)" gradientUnits="userSpaceOnUse"><stop offset=".553" stop-color="#974668" stop-opacity="0"/><stop offset=".904" stop-color="#984663"/></radialGradient></defs></g>`),
		g.Group(children),
	)
}