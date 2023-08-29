package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeInSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 7C15.434 7 2 24.271 2 30.232s13.434 23.232 30 23.232c4.615 0 8.985-1.343 12.892-3.405L55.58 57V41.686c4.017-4.4 6.42-8.935 6.42-11.454C62 24.271 48.568 7 32 7zm0 38.384c-8.283 0-15-6.784-15-15.151c0-8.369 6.717-15.151 15-15.151c8.285 0 15 6.782 15 15.151c0 8.367-6.715 15.151-15 15.151z"/><path fill="currentColor" d="M32 21.142c-.975 0-1.911.161-2.79.45A5.038 5.038 0 0 1 31 25.435c0 2.788-2.238 5.05-5 5.05a4.925 4.925 0 0 1-2.963-1.002a9.755 9.755 0 0 0-.037.749c0 5.02 4.029 9.091 9 9.091s9-4.071 9-9.091c0-5.021-4.029-9.09-9-9.09"/>`),
		g.Group(children),
	)
}