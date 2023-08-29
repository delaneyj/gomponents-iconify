package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagGermany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1FlagForFlagGermany0" d="M100.96 10.37c-1.56 0-3.08.61-4.2 1.74c-9.56 9.56-17.94 11.37-30.07 11.37c-3.68 0-7.72-.18-12-.37c-3.37-.15-6.85-.31-10.61-.4c-.66-.02-1.3-.02-1.94-.02c-30.68 0-40.49 18.56-40.89 19.35a5.946 5.946 0 0 0-.62 3.29l6.72 66.95a5.958 5.958 0 0 0 3.75 4.95c.7.27 1.43.4 2.16.4c1.43 0 2.84-.51 3.95-1.5c.1-.09 12.42-10.63 32.12-10.63c2.52 0 5.09.17 7.63.51c9.26 1.23 16.04 1.78 21.95 1.78c18.93 0 32.93-6.1 46.82-20.38a5.929 5.929 0 0 0 1.42-5.88l-20.51-66.94c-.61-2-2.22-3.53-4.25-4.03c-.47-.13-.95-.19-1.43-.19z"/></defs><use fill="#2f2f2f" href="#notoV1FlagForFlagGermany0"/><clipPath id="notoV1FlagForFlagGermany1"><use href="#notoV1FlagForFlagGermany0"/></clipPath><g clip-path="url(#notoV1FlagForFlagGermany1)"><defs><path id="notoV1FlagForFlagGermany2" d="M60.11 110.71c26.61 2.07 44.56 6.93 69.48-21.69l-9.36-37.68C97.91 76.41 83.88 76.46 53.61 73.46C28 70.93 10.24 87.41 3.76 93.19l4.23 37.17c10.51-8.74 26.32-21.66 52.12-19.65z"/></defs><use fill="#ffe000" href="#notoV1FlagForFlagGermany2"/></g><g clip-path="url(#notoV1FlagForFlagGermany1)"><defs><path id="notoV1FlagForFlagGermany3" d="M54.6 77.18c25.85 2.07 43.29 6.93 67.49-21.7l-7.93-29.24C92.47 51.32 78.85 51.37 49.44 48.38C24.56 45.84 7.3 62.33 1.01 68.11l2.96 28.72c10.2-8.73 25.57-21.66 50.63-19.65z"/></defs><use fill="#f64e3d" href="#notoV1FlagForFlagGermany3"/></g>`),
		g.Group(children),
	)
}