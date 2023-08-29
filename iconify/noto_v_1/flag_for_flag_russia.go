package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagRussia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1FlagForFlagRussia0" d="M100.96 10.37c-1.56 0-3.07.61-4.2 1.74c-9.56 9.56-17.94 11.37-30.07 11.37c-3.68 0-7.72-.18-12-.37c-3.37-.15-6.85-.3-10.61-.4c-.66-.02-1.31-.02-1.95-.02c-30.67 0-40.49 18.55-40.89 19.34a5.962 5.962 0 0 0-.62 3.29l6.72 66.95a5.958 5.958 0 0 0 3.75 4.95a6 6 0 0 0 2.16.41c1.43 0 2.84-.52 3.95-1.51c.1-.09 12.43-10.63 32.12-10.63c2.52 0 5.09.17 7.64.51c9.27 1.23 16.03 1.78 21.95 1.78c18.93 0 32.93-6.1 46.82-20.38a5.957 5.957 0 0 0 1.42-5.88l-20.51-66.95a5.963 5.963 0 0 0-4.24-4.03c-.48-.11-.96-.17-1.44-.17z"/></defs><use fill="#fff" href="#notoV1FlagForFlagRussia0"/><clipPath id="notoV1FlagForFlagRussia1"><use href="#notoV1FlagForFlagRussia0"/></clipPath><g clip-path="url(#notoV1FlagForFlagRussia1)"><defs><path id="notoV1FlagForFlagRussia2" d="M60.11 110.7c26.61 2.07 44.56 6.93 69.47-21.7l-9.36-37.68C97.9 76.4 83.88 76.45 53.61 73.46C28 70.93 10.23 87.41 3.75 93.2l4.23 37.17c10.52-8.75 26.34-21.68 52.13-19.67z"/></defs><use fill="#e74600" href="#notoV1FlagForFlagRussia2"/></g><g clip-path="url(#notoV1FlagForFlagRussia1)"><defs><path id="notoV1FlagForFlagRussia3" d="M54.6 77.18c25.85 2.08 43.29 6.93 67.5-21.7l-7.94-29.23C92.48 51.33 78.85 51.38 49.44 48.38C24.56 45.84 7.3 62.33 1.01 68.11l2.95 28.72C14.17 88.1 29.54 75.17 54.6 77.18z"/></defs><use fill="#00538b" href="#notoV1FlagForFlagRussia3"/></g>`),
		g.Group(children),
	)
}