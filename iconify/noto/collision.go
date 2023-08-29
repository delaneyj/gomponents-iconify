package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collision(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#FFA000" stroke="#F44336" stroke-miterlimit="10" stroke-width="3" d="m68.27 8.51l15.84 33.16l35.33-31.88l-18.06 45.64l19.67-5.31l-18.65 26.51l17.03 5.87l-17.4 8.27l16.93 29.03l-32.19-19.91l-6.9 21.54L69.65 98.4l-29 21.78l8.82-24.85l-39.97 8.04L39.42 81.1L7.89 58.4l35.96 3.11L10.6 10.52L64 45.51z"/><radialGradient id="notoCollision0" cx="77.587" cy="75.735" r="26.365" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#FFFDE7"/><stop offset="1" stop-color="#FFF176"/></radialGradient><path fill="url(#notoCollision0)" stroke="#FFEB3B" stroke-miterlimit="10" stroke-width="2" d="m72.63 34.77l9.54 24.71l18.88-21.31l-14.19 29.74l17.21-5.17l-15.95 15.85l13.36 4.39l-12.8 1.89l10.19 14.44l-16.56-10.58l-3.81 12.32l-5.87-14.09l-14.65 11.16L64 84.45l-23 5.02l16.95-10.28l-22.11-10.63l24.49 2.28l-20.11-31.55L71.1 60.17z"/>`),
		g.Group(children),
	)
}