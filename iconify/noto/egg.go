package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#F2A05B" d="M63.43 4.51c31.68-.07 51.67 51.46 43.78 84.12c-6.98 28.92-27.03 34.91-43.08 34.91s-35.54-7.78-43.08-33.79C12.33 59.63 33.87 4.58 63.43 4.51z"/><path fill="#FECC88" d="M39.26 29.03c-3.8 6.07-7.63 13.71-8.45 20.08c-1.5 11.73 7.13 13.05 10.14 7.23s6.09-11.71 10.98-19.52c6.76-10.79 17.83-18.02 12.39-22.71s-17.08 2.16-25.06 14.92z"/><ellipse cx="91.4" cy="43.58" fill="#D47858" rx="2.02" ry="2.06"/><ellipse cx="91.66" cy="94.05" fill="#D47858" rx="2.19" ry="2.18" transform="rotate(-67.642 91.663 94.057)"/><circle cx="32.91" cy="102.87" r="2.04" fill="#D47858" transform="rotate(-67.642 32.913 102.875)"/><ellipse cx="97.65" cy="62.57" fill="#D47858" rx="3.07" ry="2.7" transform="rotate(-27.393 97.64 62.565)"/><ellipse cx="68.52" cy="89.41" fill="#D47858" rx="3.07" ry="2.7" transform="rotate(-27.393 68.515 89.405)"/><ellipse cx="64.79" cy="104.32" fill="#D47858" rx="3.23" ry="2.7" transform="rotate(-51.428 64.796 104.32)"/><ellipse cx="52.39" cy="107.01" fill="#D47858" rx="3.58" ry="3.46" transform="rotate(-48.356 52.395 107.012)"/>`),
		g.Group(children),
	)
}