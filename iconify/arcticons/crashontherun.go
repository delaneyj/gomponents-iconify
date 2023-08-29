package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crashontherun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.55 34.77S31.86 39 24.3 37.92S14.53 25.66 11.79 27c-2.3 1.14-.42 14.48 8.44 14.48a6.27 6.27 0 0 0 5.47-3.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.83 32.14c-1.88 0-4.21.07-7.06.07c-9.36 0-8.14-8.27-12.42-8.27c-5.5 0-4.45 7.64-3.49 11.25s3.32 8.31 9.57 8.31c4.4 0 5.34-2.72 5.5-4.33m13-9.69A2.74 2.74 0 0 1 41 32a3.58 3.58 0 0 1-3.46 3.14c-2.29 0-4.66-1-4.66-2.37s1.21-3.29 5.05-3.29Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.53 16.35A28.24 28.24 0 0 0 6.55 4.5v11.84c0 3.9 2.26 5.33 4.74 5.58a5.11 5.11 0 0 0-1.76 4.74m7.78-3.37c.78-2.27 5.15-6.78 8.15-4.89s-.69 8.6-.69 8.6Zm9.43 4.59A3.07 3.07 0 0 1 31 26a3 3 0 0 1 1.1 4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.77 27a13.77 13.77 0 0 0 2 .88M32.13 30c.95.44 1.6.78 1.6.78M37 29.52c1.36-1.07 4.35-4.6 4.41-6.74a11.94 11.94 0 0 0-5.26.36"/><circle cx="22.11" cy="22.77" r="1.83" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29 28.75a1.69 1.69 0 0 0 .38-3.08m-2.56 1.9A1.38 1.38 0 0 0 27 28m-9.69-10.06c.07-4.32 4.06-6.72 7-6.72s5.87 3.9 4.61 9.27c0 0-1.05-4.78-4.66-5.2s-6.95 2.65-6.95 2.65Zm10.06 7.42s1.84-5.83 6.36-5.2s1.45 6.3 1.45 6.3a2.69 2.69 0 0 0-2.67-1.87a27.2 27.2 0 0 0-5.14.77Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.55 11.27c-.29-1.08-1.72-1.56-3.27-1.65c0 0 3.57-1.42 5.47 1c0 0 1.12-3.91 0-5c0 0 3.67 1.41 3.21 5.08a6.19 6.19 0 0 0 .31-5.5s5.53 2.14 4.46 6.42a11.8 11.8 0 0 0 .61-3.21s1.21 5.61-1.83 8.25c0 0 3-1 4.13-3.42c0 0-1.28 4.35-3.71 6a1.46 1.46 0 0 1 .32.88m-18.07-.94C15 17.77 11 11.52 9.77 10.57c.08-.15-1.56 4.65-.23 7.11a6.73 6.73 0 0 0 3.82 2.81"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.91 18.64c-.73.48-3.31 2.58-3.31 2.58"/>`),
		g.Group(children),
	)
}