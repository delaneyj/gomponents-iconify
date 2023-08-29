package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManWithChineseCapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#614b44" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 4.51A14.59 14.59 0 0 0 9.41 19.1v7.64h29.18V19.1A14.59 14.59 0 0 0 24 4.51Z"/><path fill="#ffcebf" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.15 25.63a2.73 2.73 0 0 0-2.07-2.3l-.87-.24A3 3 0 0 1 35 20.16v-3.07a2.77 2.77 0 0 0-2.08-2.68A23 23 0 0 1 24 16a23 23 0 0 1-8.9-1.61a2.77 2.77 0 0 0-2.1 2.7v3.07a3 3 0 0 1-2.23 2.93l-.87.24a2.73 2.73 0 0 0-2.07 2.3a2.69 2.69 0 0 0 2.68 2.92h.17a13.45 13.45 0 0 0 26.6 0h.17a2.69 2.69 0 0 0 2.7-2.92Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.93 25.31a1.15 1.15 0 1 0 1.14-1.15a1.15 1.15 0 0 0-1.14 1.15Zm16.14 0a1.15 1.15 0 1 1-1.14-1.15a1.15 1.15 0 0 1 1.14 1.15Z"/><path fill="#ff6242" d="M20.31 32.59a.61.61 0 0 0-.44.2a.57.57 0 0 0-.13.47a4.32 4.32 0 0 0 8.52 0a.57.57 0 0 0-.13-.47a.61.61 0 0 0-.44-.2Z"/><path fill="#ffa694" d="M24 34.37a5.3 5.3 0 0 0-3.29 1a4.38 4.38 0 0 0 6.58 0a5.3 5.3 0 0 0-3.29-1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M20.31 32.59a.61.61 0 0 0-.44.2a.57.57 0 0 0-.13.47a4.32 4.32 0 0 0 8.52 0a.57.57 0 0 0-.13-.47a.61.61 0 0 0-.44-.2Z"/><path fill="#ffb59e" d="M13.69 30.19a1.49.89 0 1 0 2.98 0a1.49.89 0 1 0-2.98 0Zm17.64 0a1.49.89 0 1 0 2.98 0a1.49.89 0 1 0-2.98 0Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 13.74c6.69 0 12.43.89 14.91 2.15a15 15 0 0 0-29.82 0c2.48-1.26 8.22-2.15 14.91-2.15Z"/><path fill="#45413c" d="M11.5 44.5a12.5 1.5 0 1 0 25 0a12.5 1.5 0 1 0-25 0Z" opacity=".15"/><path fill="#9ceb60" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5c-4.9 3.71-7.43 9.89-8 16h16c-.57-6.11-3.1-12.29-8-16Z"/><path fill="#656769" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 13.74c-9 0-16.23 1.59-16.23 3.55v3.6a2.12 2.12 0 0 0 1.64 2v-.8A1.12 1.12 0 0 1 9.94 21C13 19.15 20.24 18.79 24 18.79s11 .36 14.06 2.21a1.12 1.12 0 0 1 .53 1.12v.8a2.12 2.12 0 0 0 1.64-2v-3.6c0-1.99-7.23-3.58-16.23-3.58Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.48 16.26a2.52 2.52 0 1 0 5.04 0a2.52 2.52 0 1 0-5.04 0Z"/>`),
		g.Group(children),
	)
}