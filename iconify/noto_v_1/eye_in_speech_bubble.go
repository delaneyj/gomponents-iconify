package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeInSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#78a3ad" d="M64 19.39c-29.26 0-52.43 25.97-61.74 38.24a10.366 10.366 0 0 0 0 12.55c9.31 12.27 32.48 38.24 61.74 38.24s52.43-25.97 61.74-38.24c2.82-3.71 2.82-8.83 0-12.55C116.43 45.37 93.26 19.39 64 19.39z"/><circle cx="64" cy="63.91" r="28.3" fill="#fff"/><circle cx="64" cy="63.91" r="16.05" fill="#64878e"/><path fill="#78a3ad" d="m86.73 102.79l22.12 10.56c2.53 1.21 5.24-1.29 4.24-3.91l-8.4-22.01l-17.96 15.36z"/>`),
		g.Group(children),
	)
}