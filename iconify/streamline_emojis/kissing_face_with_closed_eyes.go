package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFaceWithClosedEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M1.5 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M21.5 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 21.5 38.5Z"/><path fill="#fff48c" d="M15.5 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#ebcb00" d="M39.39 30.44A20 20 0 0 1 31.2 39l-1.12-2.62l-.55-1.29a7.39 7.39 0 0 1 .16-5.86a5.71 5.71 0 0 1 3.9-3.32a4.46 4.46 0 0 1 4.62 1.69a4.89 4.89 0 0 1 .79 1.47Z"/><path fill="#45413c" d="M5.5 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M1.5 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ff6242" d="M46 29.88a4.64 4.64 0 0 1 .06 4.12a6 6 0 0 1-3.45 3.22L34.36 40L31 32a6 6 0 0 1 .12-4.72a4.65 4.65 0 0 1 3.15-2.68a3.65 3.65 0 0 1 4.34 2.57l.64 2.06l2-.89A3.65 3.65 0 0 1 46 29.88Z"/><path fill="#ff866e" d="M32.51 28.41A3.64 3.64 0 0 1 36.85 31l.64 2l2-.89a3.64 3.64 0 0 1 4.8 1.55a4.3 4.3 0 0 1 .52 2.15A6.27 6.27 0 0 0 46.06 34a4.64 4.64 0 0 0-.06-4.12a3.65 3.65 0 0 0-4.8-1.55l-2 .89l-.64-2.06a3.65 3.65 0 0 0-4.34-2.57a4.65 4.65 0 0 0-3.15 2.68a6 6 0 0 0-.52 2.19a4.22 4.22 0 0 1 1.96-1.05Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M46 29.88a4.64 4.64 0 0 1 .06 4.12a6 6 0 0 1-3.45 3.22L34.36 40L31 32a6 6 0 0 1 .12-4.72a4.65 4.65 0 0 1 3.15-2.68a3.65 3.65 0 0 1 4.34 2.57l.64 2.06l2-.89A3.65 3.65 0 0 1 46 29.88Zm-24.49-1.61a2 2 0 0 1 1-.27a2 2 0 1 1 0 4a2 2 0 1 1 0 4a2 2 0 0 1-1-.27"/><path fill="#ffaa54" d="M7 26.5c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S10.88 25 9.5 25S7 25.67 7 26.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.5 20.5a1.75 1.75 0 0 1-3.5 0m-21 0a1.75 1.75 0 0 1-3.5 0"/>`),
		g.Group(children),
	)
}