package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithRollingEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fcc21b" d="M63.8 9.5c-62.5 0-64 70.1-64 83.9c0 13.9 28.7 25.1 64 25.1s64-11.2 64-25.1c0-13.8-1.5-83.9-64-83.9z"/><path fill="#2f2f2f" d="M75.6 94.8h-24c-1.6 0-2.9-1.5-2.9-3.3c0-1.8 1.3-3.3 2.9-3.3h24c1.6 0 2.9 1.5 2.9 3.3c0 1.8-1.3 3.3-2.9 3.3z"/><ellipse cx="87.4" cy="58.8" fill="#fff" rx="16.6" ry="16.3"/><defs><ellipse id="notoV1FaceWithRollingEyes0" cx="87.4" cy="58.8" rx="16.6" ry="16.3"/></defs><clipPath id="notoV1FaceWithRollingEyes1"><use href="#notoV1FaceWithRollingEyes0"/></clipPath><circle cx="83.5" cy="47" r="7.9" fill="#2f2f2f" clip-path="url(#notoV1FaceWithRollingEyes1)"/><ellipse cx="40.6" cy="58.8" fill="#fff" rx="16.6" ry="16.3"/><defs><ellipse id="notoV1FaceWithRollingEyes2" cx="40.6" cy="58.8" rx="16.6" ry="16.3"/></defs><clipPath id="notoV1FaceWithRollingEyes3"><use href="#notoV1FaceWithRollingEyes2"/></clipPath><circle cx="36.7" cy="47" r="7.9" fill="#2f2f2f" clip-path="url(#notoV1FaceWithRollingEyes3)"/><path fill="#fcc21b" d="M23.6 72h34.2v8.7H23.6zm46.6 0h34.2v8.7H70.2z"/>`),
		g.Group(children),
	)
}