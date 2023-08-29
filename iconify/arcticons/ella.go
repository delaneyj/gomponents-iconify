package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31 10.91c5 0 7.9.68 10.79 4.16c2.35 2.84 2.45 9.54 1.57 11.7c-2.3 5.59-7.94 4.46-7.94 1.42M17 10.91c-5 0-7.91.68-10.79 4.16c-2.35 2.84-2.46 9.54-1.57 11.7c2.29 5.59 7.94 4.46 7.94 1.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32 11.73A6.6 6.6 0 0 0 29.28 10a16.64 16.64 0 0 0-10.56 0A6.6 6.6 0 0 0 16 11.73c-2.7 2.63-4.18 7.45-4.31 14a3.77 3.77 0 0 0 .57 2.06a26.93 26.93 0 0 0 1.6 2.3c.89 1.21 1.6 2.16 1.6 3c0 2.63 2.1 5.71 8.51 5.74s8.51-3.11 8.51-5.74c0-.87.71-1.82 1.6-3a26.93 26.93 0 0 0 1.6-2.3a3.77 3.77 0 0 0 .57-2.06c-.1-6.55-1.58-11.37-4.25-14Zm-14.16 8.15c1 0 1.88 1.33 1.88 3s-.84 3-1.88 3S16 24.48 16 22.84s.8-2.96 1.84-2.96ZM24 34.05a4.53 4.53 0 0 1-4.79-3.65a1.83 1.83 0 0 1 .31-1.58c1.27-1.62 4.35-1.55 4.48-1.56s3.21-.06 4.48 1.56a1.83 1.83 0 0 1 .31 1.58c-.23.93-.72 3.65-4.79 3.65m6.16-8.25c-1 0-1.88-1.32-1.88-3s.84-3 1.88-3s1.88 1.33 1.88 3s-.84 3-1.88 3Z"/><circle cx="17.73" cy="22.12" r=".75" fill="currentColor"/><circle cx="30.04" cy="22.12" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.35 32.44c.17 1.56 1.13 3.4 2.86 3.71a3.17 3.17 0 0 0 3.79-2.1m6.65-1.61c-.17 1.56-1.13 3.4-2.86 3.71a3.17 3.17 0 0 1-3.79-2.1"/>`),
		g.Group(children),
	)
}