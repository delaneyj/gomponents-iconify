package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSmoking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M10.83 45.28a13.17 1.72 0 1 0 26.34 0a13.17 1.72 0 1 0-26.34 0Z" opacity=".15"/><path fill="#00b8f0" d="M5.86 5.86h36.29v36.29H5.86Z"/><path fill="#4acfff" d="M42.14 11.58a5.72 5.72 0 0 0-5.72-5.72H11.58a5.72 5.72 0 0 0-5.72 5.72v3.73a5.72 5.72 0 0 1 5.72-5.72h24.84a5.72 5.72 0 0 1 5.72 5.72Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.86 5.86h36.29v36.29H5.86Z"/><path fill="#fff" d="M12 21.64h19.56v4.72H12a1.14 1.14 0 0 1-1.14-1.14v-2.44A1.14 1.14 0 0 1 12 21.64Zm22.09 0h1.14a1.14 1.14 0 0 1 1.14 1.14v2.43a1.14 1.14 0 0 1-1.14 1.14h-1.14v-4.71Z"/><path fill="#e0e0e0" d="M12.05 21.64a1.15 1.15 0 0 0-1.15 1.14v1.75a1.15 1.15 0 0 1 1.15-1.14h19.51v-1.75Zm23.18 0h-1.14v1.75h1.14a1.14 1.14 0 0 1 1.15 1.14v-1.75a1.14 1.14 0 0 0-1.15-1.14Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12 21.64h19.56v4.72h0H12a1.14 1.14 0 0 1-1.14-1.14v-2.44A1.14 1.14 0 0 1 12 21.64Zm22.09 0h1.14a1.14 1.14 0 0 1 1.14 1.14v2.43a1.14 1.14 0 0 1-1.14 1.14h-1.14h0v-4.71h0Z"/><path fill="#ff6242" d="M9.6 9.6a2.37 2.37 0 0 0 0 3.34L35.06 38.4a2.36 2.36 0 1 0 3.34-3.34L12.94 9.6a2.37 2.37 0 0 0-3.34 0Z"/><path fill="#e04122" d="M9.36 11.69a2.35 2.35 0 0 1 3.34 0l25.45 25.45a2.44 2.44 0 0 1 .56.88a2.38 2.38 0 0 0-.31-3L12.94 9.6a2.36 2.36 0 0 0-3.89 2.47a1.91 1.91 0 0 1 .31-.38Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.6 9.6a2.37 2.37 0 0 0 0 3.34L35.06 38.4a2.36 2.36 0 1 0 3.34-3.34L12.94 9.6a2.37 2.37 0 0 0-3.34 0Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.59 18.23a1.29 1.29 0 1 0 2.58 0a1.29 1.29 0 1 0-2.58 0Z"/><path fill="#fff" d="M32.23 13.81a1.09 1.09 0 0 0 2.17.19a2 2 0 0 0 .75.14a2.27 2.27 0 1 0-2.08-1.36a1.08 1.08 0 0 0-.84 1.03Z"/><path fill="#e0e0e0" d="M35.15 11.26a2.26 2.26 0 0 1 2.1 1.43a2.32 2.32 0 0 0 .16-.84a2.27 2.27 0 0 0-4.53 0a2.09 2.09 0 0 0 .17.83a2.25 2.25 0 0 1 2.1-1.42Zm-2.5 3.39a1.06 1.06 0 0 1 .42-.22a2.28 2.28 0 0 1-.19-.91a2.3 2.3 0 0 1 .13-.74a1.1 1.1 0 0 0-.78 1a1.08 1.08 0 0 0 .42.87Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.23 13.81a1.09 1.09 0 0 0 2.17.19a2 2 0 0 0 .75.14a2.27 2.27 0 1 0-2.08-1.36a1.08 1.08 0 0 0-.84 1.03Z"/>`),
		g.Group(children),
	)
}