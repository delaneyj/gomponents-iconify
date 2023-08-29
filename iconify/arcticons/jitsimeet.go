package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jitsimeet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19 34.56s.93-2-4.42-2.93c-2-.35-3.33 1.35-3.66 3.41c-.56 3.46.85 8.46.85 8.46c.39-1.32 1.58-5.46 4.29-6.11c0 0 1.64 2 5.46 2s8.25-3 9.08-5l1.44.52a11.4 11.4 0 0 0 4.26-6.27c.83-3.91 2.33-10.91-1.56-11.3s-7.54 4-7.54 4l-2.51-1.64s-1-5.87.17-7s5.91-2.75 6.53-4.51s-.24-4 .15-3.64a4.57 4.57 0 0 1 1.27 5a6.9 6.9 0 0 1-2.54 3a42.12 42.12 0 0 0 1.12-4.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.27 17.42c-.13 4.54 0 6 .44 6.93s2.29 5.2 2.29 5.2s-6-4.38-9.4-4.77a9.54 9.54 0 0 0-6.29 1.94c2.47-.4 3.73-.74 5.17 0s3.25 3.23 5.38 3.4a19.58 19.58 0 0 0 5.14-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.69 19.7c-1.79 0-5.66 0-7.3 3.44c-1.32 2.75.6 6.85 3.31 8.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.54 31.63c0 1.08-.93 4.74 1.48 5.76c.89.37 2.95-2.63 3.49-3.39a14.67 14.67 0 0 0 1.19-2.69c1.71.12 7.31-2.57 7.31-2.57m2.55 5.67a12 12 0 0 0 .8-4.24m-1.09-17.65c.6-.18 1.42.5 1.85 1.9s.4 3-.57 3.53s-2.58-.79-2.71-2.07s.46-3.07 1.43-3.36Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.31 16.17A14.88 14.88 0 0 1 29 14.42"/>`),
		g.Group(children),
	)
}