package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagFrance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1FlagForFlagFrance0" d="M100.96 10.37c-1.56 0-3.08.61-4.2 1.74c-9.56 9.56-17.94 11.37-30.07 11.37c-3.68 0-7.72-.18-11.99-.37c-3.37-.15-6.85-.3-10.62-.4c-.66-.02-1.31-.02-1.95-.02c-30.67 0-40.48 18.56-40.89 19.35a5.946 5.946 0 0 0-.62 3.29l6.72 66.95a5.966 5.966 0 0 0 3.75 4.94c.7.28 1.43.4 2.16.4c1.43 0 2.84-.52 3.95-1.5c.1-.09 12.43-10.63 32.13-10.63c2.52 0 5.09.17 7.63.51c9.27 1.23 16.04 1.78 21.95 1.78c18.93 0 32.93-6.1 46.81-20.38a5.941 5.941 0 0 0 1.42-5.89l-20.51-66.95a5.96 5.96 0 0 0-4.25-4.03c-.46-.1-.94-.16-1.42-.16z"/></defs><use fill="#f87055" href="#notoV1FlagForFlagFrance0"/><clipPath id="notoV1FlagForFlagFrance1"><use href="#notoV1FlagForFlagFrance0"/></clipPath><path fill="#00538b" d="m36.12 20.85l12.11 87.1l-31.7 11.86l-9.53.51l-10.05-79.88l23.19-17.53z" clip-path="url(#notoV1FlagForFlagFrance1)"/><g clip-path="url(#notoV1FlagForFlagFrance1)"><path fill="#fff" d="M39.94 14.91c-2.44-.08-4.72.01-6.89.19c2.68 20.27 8.17 58.93 15.45 95.36c3.59-.17 7.4-.03 11.44.61c12.85 2.03 23.58 2.99 33.16 1.35c-9.8-35.78-18.08-75.13-22.38-96.73c-8.7.72-18.62-.4-30.78-.78z"/></g>`),
		g.Group(children),
	)
}