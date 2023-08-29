package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-digit-ninea" gradientUnits="userSpaceOnUse" x1="64.005" y1="101.667" x2="64.005" y2="24.97" gradientTransform="matrix(1 0 0 -1 0 128)"><stop offset="0" stop-color="#757575"/><stop offset=".515" stop-color="#504f4f"/></linearGradient><path d="M75.34 74.48c.24-1.31-1.29-2.26-2.34-1.45c-3.5 2.7-7.43 4.05-11.79 4.05c-6.56 0-11.79-2.3-15.69-6.9s-5.85-10.64-5.85-18.13c0-4.92.99-9.38 2.97-13.38s4.8-7.11 8.46-9.33C54.76 27.11 59 26 63.82 26c7.52 0 13.49 2.8 17.9 8.41s6.62 13.11 6.62 22.51v3.49c0 13.47-3.04 23.74-9.13 30.82S64.04 101.9 51.97 102c-.82 0-1.48-.66-1.48-1.48v-7.34c0-.82.66-1.48 1.48-1.48h.31c7.35-.1 12.99-1.87 16.92-5.31c3.16-2.77 5.21-6.73 6.14-11.91zM64.09 67.3c2.37-.05 4.6-.76 6.69-2.13c.02-.01.04-.03.06-.04c3.18-2.11 4.96-5.78 4.96-9.6v-1.08c0-5.45-1.12-9.83-3.35-13.15s-5.13-4.98-8.68-4.98s-6.39 1.46-8.51 4.39s-3.17 6.63-3.17 11.1c0 4.7 1.1 8.48 3.29 11.34c2.05 2.68 5.34 4.22 8.71 4.15z" fill="url(#ssvg-id-digit-ninea)"/>`),
		g.Group(children),
	)
}