package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f1367id0)" d="M2.195 16.002c0-1.69 1.38-3.06 3.06-3.06h21.85c1.69 0 3.06 1.38 3.06 3.06c0 1.69-1.38 3.06-3.06 3.06H5.255c-1.68 0-3.06-1.38-3.06-3.06Z"/><path fill="url(#f1367id1)" d="M2.195 16.002c0-1.69 1.38-3.06 3.06-3.06h21.85c1.69 0 3.06 1.38 3.06 3.06c0 1.69-1.38 3.06-3.06 3.06H5.255c-1.68 0-3.06-1.38-3.06-3.06Z"/><path fill="url(#f1367id2)" d="M2.195 16.002c0-1.69 1.38-3.06 3.06-3.06h21.85c1.69 0 3.06 1.38 3.06 3.06c0 1.69-1.38 3.06-3.06 3.06H5.255c-1.68 0-3.06-1.38-3.06-3.06Z"/><g filter="url(#f1367id3)"><path fill="#544C5D" d="M5.441 14.487h23.366v2.276H5.441z"/></g><defs><linearGradient id="f1367id0" x1="16.18" x2="16.18" y1="12.941" y2="19.061" gradientUnits="userSpaceOnUse"><stop stop-color="#4A4253"/><stop offset="1" stop-color="#39204E"/></linearGradient><linearGradient id="f1367id1" x1="2.195" x2="5.441" y1="16.002" y2="16.002" gradientUnits="userSpaceOnUse"><stop offset=".314" stop-color="#342D3C"/><stop offset="1" stop-color="#342D3C" stop-opacity="0"/></linearGradient><radialGradient id="f1367id2" cx="0" cy="0" r="1" gradientTransform="matrix(-3.33793 0 0 -3.68387 29.768 14.942)" gradientUnits="userSpaceOnUse"><stop stop-color="#6C6673"/><stop offset="1" stop-color="#6C6673" stop-opacity="0"/></radialGradient><filter id="f1367id3" width="26.366" height="5.276" x="3.941" y="12.987" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feGaussianBlur result="effect1_foregroundBlur_18590_1731" stdDeviation=".75"/></filter></defs></g>`),
		g.Group(children),
	)
}