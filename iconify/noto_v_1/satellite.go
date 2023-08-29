package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#006ca0" d="m80.3 56.8l44.9 23.8c.7.3 1.1 1 1.1 1.8v25.5c0 1.5-1.6 2.5-2.9 1.8L80.6 87.1c-.7-.4-1.1-1.1-1.1-1.8l.8-28.5z"/><path fill="#40c0e7" d="M125.4 80.8L97.5 96l-11.2-5.9l28.9-14.8l10.2 5.4v.1z"/><defs><path id="notoV1Satellite0" d="m80.3 56.8l44.9 23.8c.7.3 1.1 1 1.1 1.8v25.5c0 1.5-1.6 2.5-2.9 1.8L80.7 87.1c-.7-.4-1.1-1.2-1.1-2l.7-28.3z"/></defs><clipPath id="notoV1Satellite1"><use href="#notoV1Satellite0"/></clipPath><g clip-path="url(#notoV1Satellite1)"><path fill="#b8e9f4" d="m80.563 67.42l.939-1.765l46.524 24.753l-.94 1.766z"/></g><g clip-path="url(#notoV1Satellite1)"><path fill="#b8e9f4" d="m78.88 77.555l.932-1.77l48.133 25.35l-.932 1.77z"/></g><path fill="#64878e" d="m66.9 41.2l-16.6 9.6c-1.2.7-1.9 1.8-2 3.2l-1.4 17.2c-.2 2.7 1.4 5.2 3.9 6.1l16.1 5.8c1.2.4 2.6.3 3.7-.5l16.1-10.9l-19.8-30.5z"/><path fill="#78a3ad" d="M67.6 41.2c-.7-.3-1.4.2-1.4.9L65.9 61c-.1 1.8.9 3.4 2.6 4L86 71.7c.7.3 1.4-.3 1.4-1c-.2-4.4-1.1-16.7-3.6-20.4c-2.2-3.2-12.5-7.6-16.2-9.1z"/><path fill="#294f6e" d="M77.4 42.4s-8.8 4.1-7.8 13.7c.8 7.5 8.3 11.1 16.7 9.1l-8.9-22.8z"/><ellipse cx="81.9" cy="53.8" fill="#006ca0" rx="9.6" ry="12.2" transform="rotate(-21.292 81.853 53.839)"/><path fill="#78a3ad" d="M65.9 61.6L47.1 73.3s.5 2.6 3 3.6L68.4 65s-1.9-1.1-2.5-3.4z"/><path fill="#294f6e" d="M75.5 46.3c-2.2 2.3-2.7 9 .3 12.8c3.3 4.2 8.5 5.5 11.9 4c1.9-.8-2.9-2-7-7.7c-3.7-4.9-2.7-11.7-5.2-9.1z"/><path fill="#fff" d="m79.6 55.2l6.6-4"/><path fill="#fcc21b" d="M79.6 56.7c-.5 0-1-.3-1.3-.7c-.4-.7-.2-1.6.5-2.1l6.6-3.9c.7-.4 1.6-.2 2.1.5s.2 1.6-.5 2.1l-6.6 3.9c-.3.1-.5.2-.8.2z"/><path fill="#006ca0" d="M44.6 40.4c0-.7-.4-1.4-1-1.8L4.1 16.8c-1.3-.8-3 .2-3 1.7v27.4c0 .7.4 1.4 1.1 1.8l39.1 20.1c1.3.7 2.9-.3 2.9-1.7l.4-25.7z"/><path fill="#40c0e7" d="m1.2 40.8l24.4-12.2L37 34.9L8.8 51.1l-6.5-3.4c-.7-.3-1.1-1-1.1-1.8v-5.1zm43.5 3L22 57.9l6.5 3.3l16-10.1z"/><defs><path id="notoV1Satellite2" d="M44.6 40.4c0-.7-.4-1.4-1-1.8L4.1 16.8c-1.3-.8-3 .2-3 1.7v27.4c0 .7.4 1.4 1.1 1.8l39.1 20.1c1.3.7 2.9-.3 2.9-1.7l.4-25.7z"/></defs><clipPath id="notoV1Satellite3"><use href="#notoV1Satellite2"/></clipPath><g fill="#b8e9f4" clip-path="url(#notoV1Satellite3)"><path d="m45 49.9l-45.3-25l.9-1.8l45.3 25.1zm0 10.5L-1 35.8l1-1.7l46 24.5z"/></g>`),
		g.Group(children),
	)
}