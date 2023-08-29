package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Howlr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.25 8.93c0-1.68-1.68-3.43-2.94-3.43s-4.81 9-4.81 9a13.67 13.67 0 0 0-3.57 8.37c0 1.25-1.76 2.81-1.76 6.09s4.29 9 4.29 9l-.41-1.42c.59.93 4.32 3.91 5 4a7.58 7.58 0 0 0-1.49.74h1l-.49 1.23m23.46-16.95s1 1.23 1.65 2.35a6.17 6.17 0 0 1 .6 3l-1-1.15a25.55 25.55 0 0 1-1.15 4.92c-1 3.11-7.88 6.15-8.76 6.32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.57 27v-3.32c0-3.02-2.82-7.17-2.82-7.17a4.75 4.75 0 0 1 3.14-1.25a2.2 2.2 0 0 1 1.29.53M23.83 9.71A14 14 0 0 0 19.05 9a11.72 11.72 0 0 0-4.28 1.57a3.71 3.71 0 0 0 3.83.55l-.75 1.49a4.19 4.19 0 0 0 2.77-.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.32 21.74c.9-.11 3.51-2 3.51-4.47s-1.08-6.74-2.35-7.06s-6.33 2.3-6.33 2.3a13 13 0 0 0-6.28-4.15c-2.18-.44-6.5-2.28-6.87-2.86a8 8 0 0 0 .87 3.55m-9.26 21.12c.53.56 2.46.44 3 .44a17.41 17.41 0 0 0 2.93-1a45.09 45.09 0 0 0 5.48.72c1.65 0 6.4-1.2 7.24-3.94a4.75 4.75 0 0 1 3.8 4.73c0 4-3.8 6.28-7.91 6.28s-8.33-2-9.89-4.23a19.66 19.66 0 0 1-1.65-2.56"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 34c2.93 0 11.37-1.14 16.26-5.69m-1.19-2.75s1.17-.36 1.17-3a6.9 6.9 0 0 0-6.15-7c-2.51 0-5 3-5 6.71s1.26 3.94 1.26 3.94M23.18 13a6.14 6.14 0 0 1 3.19-1.23a9.42 9.42 0 0 1 3.27 1m-16.93.74A3.7 3.7 0 0 1 15.13 12c1 0 1.34.31 1.34.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.73 27c.19-.33 1.68-2.2 5.24-2.2a9.39 9.39 0 0 1 5.57 1.61m-16.75-2.3a6.55 6.55 0 0 1 3.84 1.59c.15.31-2.7 3.91-3.07 3.91s-3.62-2.72-3.62-4.05s2.31-1.45 2.85-1.45Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 25.45a9.09 9.09 0 0 1 3.29-.5a7.81 7.81 0 0 1 1.79.16m9.36-7.28a3.55 3.55 0 0 1 2.38 3.28c0 1.64-1.23 2.89-2 2.89a4.09 4.09 0 0 1-1.72-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.29 18.42c-.74 0-2.35.75-2.35 3s1.23 2.81 1.23 2.81a7.84 7.84 0 0 0 1.3-4.23c0-2.46-1.21-4.26-2.53-4.26c-1 0-4.1 3.65-4.1 7.15a6.65 6.65 0 0 0 .36 2.21m4.88 15.43a12.3 12.3 0 0 0-4.54.24m25.32 1.16s-.72-1-3-1.16a22.45 22.45 0 0 0-5 .36a7.7 7.7 0 0 0-2.33 1.12"/>`),
		g.Group(children),
	)
}