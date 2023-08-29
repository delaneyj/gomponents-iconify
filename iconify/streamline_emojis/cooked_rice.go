package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookedRice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M6.13 44.86a17.87 2.14 0 1 0 35.74 0a17.87 2.14 0 1 0-35.74 0Z" opacity=".15"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.49 42.91c0 1.43-4.25 2.59-9.49 2.59s-9.49-1.16-9.49-2.59v-3.45h19Z"/><path fill="#ff866e" d="M3.3 17.89s4.31 6 20.7 6s20.7-6 20.7-6S48.16 42.05 24 42.05S3.3 17.89 3.3 17.89Z"/><path fill="#ffa694" d="M24 29.41c12.41 0 18.29-3.05 20.65-4.82a27.89 27.89 0 0 0 .05-6.7s-4.31 6-20.7 6s-20.7-6-20.7-6a27.89 27.89 0 0 0 .05 6.7c2.36 1.77 8.24 4.82 20.65 4.82Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.3 17.89s4.31 6 20.7 6s20.7-6 20.7-6S48.16 42.05 24 42.05S3.3 17.89 3.3 17.89Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.3 17.46a20.7 6.47 0 1 0 41.4 0a20.7 6.47 0 1 0-41.4 0Z"/><path fill="#f0f0f0" d="M24 3.52c-12.82 0-17.25 12.61-17.25 12.61s3.71 4.35 17.25 4.35s17.25-4.35 17.25-4.35S36.82 3.52 24 3.52Z"/><path fill="#fff" d="M6.87 16.25A20 20 0 0 1 24 6.94a20 20 0 0 1 17.13 9.31l.12-.12S36.82 3.52 24 3.52S6.75 16.13 6.75 16.13Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 3.52c-12.82 0-17.25 12.61-17.25 12.61s3.71 4.35 17.25 4.35s17.25-4.35 17.25-4.35S36.82 3.52 24 3.52Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.77 10.09c.36.7-.62 1.93-2.19 2.75s-3.15.93-3.52.23s.62-1.94 2.19-2.76s3.15-.92 3.52-.22ZM22 13.27c.06.79-1.32 1.54-3.1 1.67s-3.25-.39-3.31-1.18s1.33-1.54 3.1-1.67s3.25.4 3.31 1.18Zm11.42 3.35c-.06.79-1.55 1.32-3.32 1.18s-3.16-.88-3.1-1.67s1.55-1.31 3.32-1.18s3.16.88 3.1 1.67Zm-2.97-8.03c-.16.77-1.7 1.1-3.44.74S24 8 24.15 7.27s1.71-1.1 3.45-.74s3.02 1.29 2.85 2.06Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37 13.94c.07.78-1.31 1.55-3.08 1.71s-3.26-.35-3.33-1.14s1.3-1.55 3.07-1.71s3.26.35 3.34 1.14Z"/>`),
		g.Group(children),
	)
}