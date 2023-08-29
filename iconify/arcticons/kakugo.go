package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kakugo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.05 13c2.26-.71 7-1.79 9.08-3.33m-5.81 4.19a4.51 4.51 0 0 1 1.49 2.54c0 .61-1.57 6.3-1.9 7.83s-1.08 4.63-1.08 4.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.76 31.58c1.91-.17 8-3.54 11.23-4.06s7.79-1.58 11.51.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.17 28.44c.45-1.67 1.06-4.69 1.28-6.18s-.79-2.4-2.37-2.18s-8.42 2-8.42 2m3.45 11.78c2.79-1.31 5.47-2.67 7.37-1s.51 5.08 0 6.22s-1.13 1.35-1.13 1.35a46.45 46.45 0 0 0-5 .62m-4.01-7.25a9.31 9.31 0 0 1 2.77 5.66M11.63 35.6c2.37-1.13 3.42-2.34 5-1.8s.59 3.46.6 4.3a5.79 5.79 0 0 0 .34 1.71s-4.55 1.36-5.77 1.54m-3.26-6.46c.87 1.68 1.87 2.9 2.18 5.11M12 29.83a13.86 13.86 0 0 0 5.19-1.5M11.5 25.2a8.37 8.37 0 0 0 5.16-2.41M4.5 20.39a20.14 20.14 0 0 0 7.13-2.82a15.82 15.82 0 0 1 5.6-1.86M11 6.65a7.74 7.74 0 0 1 5.69 3.27"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.06 11.29A6.3 6.3 0 0 1 14.73 8"/>`),
		g.Group(children),
	)
}