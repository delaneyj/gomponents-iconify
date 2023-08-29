package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.5 45.5a12.5 1.5 0 1 0 25 0a12.5 1.5 0 1 0-25 0Z" opacity=".15"/><path fill="#ffcebf" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M39.17 32a2.56 2.56 0 0 0-1.94-2.15l-.82-.23a2.86 2.86 0 0 1-2.1-2.76V24a2.61 2.61 0 0 0-2-2.52A21.6 21.6 0 0 1 24 23a21.6 21.6 0 0 1-8.36-1.52a2.61 2.61 0 0 0-2 2.52v2.88a2.86 2.86 0 0 1-2.1 2.76l-.82.23a2.48 2.48 0 0 0 .58 4.9h.15a12.64 12.64 0 0 0 25 0h.15A2.53 2.53 0 0 0 39.17 32Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.65 32.4a1.08 1.08 0 1 0 1.08-1.07a1.08 1.08 0 0 0-1.08 1.07Zm14.7 0a1.08 1.08 0 1 1-1.08-1.07a1.08 1.08 0 0 1 1.08 1.07Z"/><path fill="#ff6242" d="M21 38.32a.48.48 0 0 0-.36.16a.5.5 0 0 0-.11.39a3.56 3.56 0 0 0 7 0a.5.5 0 0 0-.11-.39a.48.48 0 0 0-.36-.16Z"/><path fill="#ffa694" d="M24 39.78a4.39 4.39 0 0 0-2.71.83a3.6 3.6 0 0 0 5.42 0a4.39 4.39 0 0 0-2.71-.83Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21 38.32a.48.48 0 0 0-.36.16a.5.5 0 0 0-.11.39a3.56 3.56 0 0 0 7 0a.5.5 0 0 0-.11-.39a.48.48 0 0 0-.36-.16Z"/><path fill="#ffb59e" d="M14.78 36.29a1.4.84 0 1 0 2.8 0a1.4.84 0 1 0-2.8 0Zm15.64 0a1.4.84 0 1 0 2.8 0a1.4.84 0 1 0-2.8 0Z"/><path fill="#525252" d="M34.2 2.32A37.17 37.17 0 0 0 24 1a37.17 37.17 0 0 0-10.2 1.32A5.91 5.91 0 0 0 9.5 8v20.5a1 1 0 0 0 1 1h27a1 1 0 0 0 1-1V8a5.91 5.91 0 0 0-4.3-5.68Z"/><path fill="#656769" d="M34.2 2.32A37.17 37.17 0 0 0 24 1a37.17 37.17 0 0 0-10.2 1.32A5.91 5.91 0 0 0 9.5 8v4a5.91 5.91 0 0 1 4.3-5.68A37.17 37.17 0 0 1 24 5a37.17 37.17 0 0 1 10.2 1.32A5.91 5.91 0 0 1 38.5 12V8a5.91 5.91 0 0 0-4.3-5.68Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.2 2.32A37.17 37.17 0 0 0 24 1a37.17 37.17 0 0 0-10.2 1.32A5.91 5.91 0 0 0 9.5 8v20.5a1 1 0 0 0 1 1h27a1 1 0 0 0 1-1V8a5.91 5.91 0 0 0-4.3-5.68Z"/><path fill="#debb7e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.69 29.5V33a10.69 10.69 0 0 1-21.38 0v-3.5h-2v3.38a12.64 12.64 0 0 0 25.28 0V29.5Z"/>`),
		g.Group(children),
	)
}