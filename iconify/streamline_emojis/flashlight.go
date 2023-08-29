package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#656769" d="m25.59 32l18.93-20.74c.78-.78-.46-3.28-2.77-5.59s-4.81-3.55-5.59-2.77L15.38 21.83Z"/><path fill="#87898c" d="M38.75 6c.72-.73 2.58-.08 4.5 1.44a16.46 16.46 0 0 0-1.5-1.72c-2.31-2.31-4.81-3.55-5.59-2.77L15.38 21.83l2.94 2.94Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m25.59 32l18.93-20.74c.78-.78-.46-3.28-2.77-5.59s-4.81-3.55-5.59-2.77L15.38 21.83Z"/><path fill="#656769" d="M23.54 37.41L25.59 32c.94-.95-.56-4-3.36-6.79s-5.84-4.3-6.79-3.36l-5.55 1.91Z"/><path fill="#87898c" d="M17.88 24.46c.7-.7 2-.67 3.45-.05c-2.53-2.25-5-3.36-5.89-2.52l-5.55 1.87l3 3Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.54 37.41L25.59 32c.94-.95-.56-4-3.36-6.79s-5.84-4.3-6.79-3.36l-5.55 1.91Z"/><path fill="#45413c" d="M10.03 45.15a17.25 1.85 0 1 0 34.5 0a17.25 1.85 0 1 0-34.5 0Z" opacity=".15"/><path fill="#656769" d="m17 43.92l6.51-6.51c1.27-1.28-.75-5.37-4.52-9.14s-7.86-5.79-9.13-4.51l-6.51 6.5Z"/><path fill="#87898c" d="M12.5 26.55c1-1 3.1-.64 5.46.72c-3.45-3.11-6.96-4.67-8.07-3.51l-6.51 6.5L6.08 33Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m17 43.92l6.51-6.51c1.27-1.28-.75-5.37-4.52-9.14s-7.86-5.79-9.13-4.51l-6.51 6.5Z"/><path fill="#fffacf" d="M3.38 30.26c1.27-1.26 5.36.74 9.13 4.52s5.8 7.86 4.52 9.14s-5.36-.75-9.13-4.52s-5.79-7.86-4.52-9.14Z"/><path fill="#ffe500" d="M17.33 42.12c-1.87-.39-4.68-2.18-7.33-4.83S5.57 31.84 5.18 30a2 2 0 0 0-1.8.3c-1.27 1.28.75 5.37 4.52 9.14s7.86 5.79 9.13 4.52a1.91 1.91 0 0 0 .3-1.84Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.38 30.26c1.27-1.26 5.36.74 9.13 4.52s5.8 7.86 4.52 9.14s-5.36-.75-9.13-4.52s-5.79-7.86-4.52-9.14Z"/>`),
		g.Group(children),
	)
}