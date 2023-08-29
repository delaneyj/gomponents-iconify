package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThongSandal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="notoThongSandal0" x1="62.065" x2="62.065" y1="20.937" y2="126.198" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#B2DFDB"/><stop offset="1" stop-color="#80CBC4"/></linearGradient><path fill="url(#notoThongSandal0)" d="M94.17 48.8c0-18.11-8.13-40.74-31.87-44.3c-17.42-2.61-35.5 4.47-31.87 32.8c2.37 18.57 9.39 18.98 7.3 51.16C36.05 114.3 47.64 124 62.3 124s24.18-6.76 26.54-27.31c2.36-20.55 5.33-36.44 5.33-47.89z"/><linearGradient id="notoThongSandal1" x1="64.837" x2="64.837" y1="90.719" y2="49.546" gradientUnits="userSpaceOnUse"><stop offset=".334" stop-color="#0277BD"/><stop offset="1" stop-color="#01579B"/></linearGradient><path fill="url(#notoThongSandal1)" d="M87.56 46.58c-4.88-13.15-24.15-20.9-30.54-22.47l-3.85-5.56c-1.3-1.99-2.8-2.61-4.3-1.69c-1.72 1.05-.78 2.49-.11 4.39l1.91 6c-3.72 4.64-12.08 15.03-10.35 33.76c1.48 15.95 7.68 25 7.9 25.46c.62 1.33 3.09 2.59 4.84 1.78s1.99-3.02 1.14-4.85c-.03-.07-3.14-9.39-5.24-15.71c-1.28-3.84-1.25-8 .1-11.82c2.15-6.06 4.88-15.13 9.39-20.12c0 0 10.18 6.37 17.12 12.8c4.3 3.98 6.76 9.54 6.05 15.36c-.72 5.88-1.76 13.81-2.57 17.82c-.45 2.24-.7 4.52 1.56 5.39c2.26.87 3.56-1.04 4.41-3.24c.35-.93 8.15-22.19 2.54-37.3z"/><radialGradient id="notoThongSandal2" cx="59.012" cy="40.824" r="31.01" gradientTransform="matrix(.9795 -.2016 .1373 .6672 -4.394 25.484)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#40BDF6"/><stop offset="1" stop-color="#039BE5"/></radialGradient><path fill="url(#notoThongSandal2)" d="M87.56 46.58c-4.88-13.15-24.15-20.9-30.54-22.47l-3.85-5.56c-1.3-1.99-2.8-2.61-4.3-1.69c-1.72 1.05-.78 2.49-.11 4.39l1.91 6c-3.72 4.64-12.08 15.03-10.35 33.76c.55 5.95 3.06 14.88 3.06 14.88c-.01-5.98-.06-7.79 2.12-15.27c2.17-7.47 5.21-18.99 12.38-26.94c0 0 13.94 7.91 22.72 18.01c8.83 10.16 7.67 20.25 7.67 20.25c1.73-8.46 1.68-18.93-.71-25.36z"/>`),
		g.Group(children),
	)
}