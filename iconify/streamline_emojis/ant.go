package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.82 16.55S8.88 4.81 1.49 8.45m13.86 6.64s.32-10.35-7.53-9.4m18.6 32.21s-.25 6.3 1.8 6.3"/><path fill="#45413c" d="M5.9 44.2a18.6 1.8 0 1 0 37.2 0a18.6 1.8 0 1 0-37.2 0Z" opacity=".15"/><path fill="#ff866e" d="M5.42 22.49a8.4 8.4 0 1 0 16.8 0a8.4 8.4 0 1 0-16.8 0Z"/><path fill="#ffa694" d="M13.82 16.55a8.39 8.39 0 0 1 8.3 7.17a8.28 8.28 0 0 0 .1-1.23a8.4 8.4 0 0 0-16.8 0a8.28 8.28 0 0 0 .1 1.23a8.39 8.39 0 0 1 8.3-7.17Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.42 22.49a8.4 8.4 0 1 0 16.8 0a8.4 8.4 0 1 0-16.8 0Z"/><path fill="#ff866e" d="M16.22 30.89a4.8 4.8 0 1 0 9.6 0a4.8 4.8 0 1 0-9.6 0Z"/><path fill="#ffa694" d="M21 28.64a4.79 4.79 0 0 1 4.6 3.52a4.49 4.49 0 0 0 .2-1.27a4.8 4.8 0 0 0-9.6 0a4.88 4.88 0 0 0 .19 1.27A4.79 4.79 0 0 1 21 28.64Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.22 30.89a4.8 4.8 0 1 0 9.6 0a4.8 4.8 0 1 0-9.6 0Z"/><path fill="#ff866e" d="M35.13 23.1c6.46.6 11.43 7.48 11.1 10.61s-6.41 9-12.87 8.44s-11-5.35-10.51-10.61s5.82-9.04 12.28-8.44Z"/><path fill="#ffa694" d="M35.13 26.6c5.19.48 9.42 5 10.71 8.42a4.37 4.37 0 0 0 .39-1.31c.33-3.13-4.65-10-11.1-10.61s-11.79 3.18-12.28 8.44A8.26 8.26 0 0 0 23 34c1-4.69 6.09-8 12.13-7.4Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.13 23.1c6.46.6 11.43 7.48 11.1 10.61s-6.41 9-12.87 8.44s-11-5.35-10.51-10.61s5.82-9.04 12.28-8.44Zm-17.71 8.99l-6 2a2.41 2.41 0 0 0-1.56 1.66c-.58 2.41-2.12 7.48-3.86 7.48m13.22-10.17l-2.64 2.13a2.36 2.36 0 0 0-.89 2c.11 2.39.19 7.52-1.27 7.52"/><path fill="#45413c" d="M11.52 20.09a1.2 1.2 0 1 1-1.2-1.2a1.2 1.2 0 0 1 1.2 1.2Z"/><path fill="#ffcabf" d="M12.62 22.92c0 .33.54.6 1.2.6s1.2-.27 1.2-.6s-.54-.6-1.2-.6s-1.2.27-1.2.6Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.46 23.24s-2.64 1.86-3.91 6.64c-1.14 4.3-.71 9.26 1.17 11m5.81 1.31a14.71 14.71 0 0 1-1.47-10.13c1.47-6.78 5.2-8.17 5.2-8.17m3.15 15.93s-1.93-.64-2-4.8a8.19 8.19 0 0 1 4-7.32"/>`),
		g.Group(children),
	)
}