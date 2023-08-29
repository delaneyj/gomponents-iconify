package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<mask id="fluentEmojiHighContrastGreenSquare0" width="28" height="29" x="2" y="2" maskUnits="userSpaceOnUse" style="mask-type:alpha"><rect width="28" height="28" fill="#C4C4C4" rx="2" transform="matrix(0 -1 -1 0 30 30.145)"/></mask><g fill="none"><rect width="28" height="28" stroke="currentColor" stroke-width="2" rx="2" transform="matrix(-1 0 0 1 30 2)"/><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" mask="url(#fluentEmojiHighContrastGreenSquare0)"><path d="M30.31-20.958L1.69 7.662m28.62-13.62L1.69 22.662m28.62-13.62L1.69 37.662m28.62-58.62L1.69 7.662m28.62-13.62L1.69 22.662m28.62-13.62L1.69 37.662m28.62-53.62L1.69 12.662M30.31-.958L1.69 27.662m28.62-13.62L1.69 42.662m28.62-23.62L1.69 47.662m28.62-58.62L1.69 17.662m28.62-13.62L1.69 32.662m28.62-8.62L1.69 52.662"/></g><path fill="currentColor" d="M28 1H4a3 3 0 0 0-3 3v24a3 3 0 0 0 3 3h24a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3Zm1 3v.291L4.291 29H4a1 1 0 0 1-1-1v-.587L27.413 3H28a1 1 0 0 1 1 1Zm0 2.412v2.88L9.291 29H6.412L29 6.412ZM11.413 29L29 11.413v2.878L14.291 29h-2.879Zm5 0L29 16.413v2.878L19.291 29h-2.878Zm5 0L29 21.413v2.878L24.291 29h-2.878Zm5 0L29 26.413V28a1 1 0 0 1-1 1h-1.587ZM3 25.291v-2.878L22.413 3h2.878L3 25.291ZM20.291 3L3 20.291v-2.878L17.413 3h2.878Zm-5 0L3 15.291v-2.879L12.412 3h2.88Zm-5 0L3 10.291V7.412L7.412 3h2.88Zm-5 0L3 5.291V4a1 1 0 0 1 1-1h1.291Z"/></g>`),
		g.Group(children),
	)
}