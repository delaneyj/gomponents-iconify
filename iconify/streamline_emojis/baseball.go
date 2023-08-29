package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baseball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M7.62 44.86a16.38 2.14 0 1 0 32.76 0a16.38 2.14 0 1 0-32.76 0Z" opacity=".15"/><path fill="#f0f0f0" d="M4.06 24a19.94 19.94 0 1 0 39.88 0a19.94 19.94 0 1 0-39.88 0Z"/><path fill="#fff" d="M24 9.76a19.94 19.94 0 0 1 19.72 17.09a20.55 20.55 0 0 0 .22-2.85a19.94 19.94 0 0 0-39.88 0a20.55 20.55 0 0 0 .22 2.85A19.94 19.94 0 0 1 24 9.76Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4.06 24a19.94 19.94 0 1 0 39.88 0a19.94 19.94 0 1 0-39.88 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29 4.69v.08s-1.43 9.73 3.56 17.09c4.18 6.18 9.36 6.5 10.93 6.45M4.88 18.33c1 .22 8.39 2.05 12.71 7.81c4.09 5.46 2.62 15.18 2.22 17.36"/><path fill="#ff866e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.59 5.68L27 5.64a1.43 1.43 0 0 0 0 2.85h3.56a1.43 1.43 0 1 0 0-2.85Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.25 13a1.43 1.43 0 0 0-1.63-1.18l-3.52.56A1.43 1.43 0 0 0 25.92 14a1.41 1.41 0 0 0 1.63 1.18l3.52-.56A1.43 1.43 0 0 0 32.25 13Zm-18.57 5.51a1.42 1.42 0 0 0-1.95.48l-1.84 3.06a1.43 1.43 0 0 0 2.45 1.46l1.83-3.05a1.42 1.42 0 0 0-.49-1.95Z"/><path fill="#ff866e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.87 15.85a1.42 1.42 0 0 0-1.77 1l-1 3.41a1.43 1.43 0 0 0 2.69.74l1-3.41a1.42 1.42 0 0 0-.92-1.74Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.28 22.9a1.42 1.42 0 0 0-2-.12l-2.66 2.36a1.42 1.42 0 1 0 1.89 2.13l2.66-2.36a1.42 1.42 0 0 0 .11-2.01Zm3.56 6.69a1.43 1.43 0 0 0-1.76-1l-3.43 1a1.43 1.43 0 1 0 .77 2.75l3.43-1a1.42 1.42 0 0 0 .99-1.75ZM22.43 35l-3.56-.08a1.43 1.43 0 1 0-.07 2.85l3.57.08a1.43 1.43 0 0 0 .06-2.85Zm-.32 5.76l-3.49-.7a1.42 1.42 0 0 0-1.68 1.11a1.44 1.44 0 0 0 1.12 1.68l3.49.7a1.42 1.42 0 0 0 1.68-1.11a1.44 1.44 0 0 0-1.12-1.68ZM33.92 18.3a1.42 1.42 0 0 0-2-.43L29 19.8a1.43 1.43 0 0 0 1.54 2.4l3-1.93a1.42 1.42 0 0 0 .38-1.97Zm3.46 4.42a1.42 1.42 0 0 0-2 .41l-2 3a1.43 1.43 0 1 0 2.38 1.57l2-3a1.42 1.42 0 0 0-.38-1.98Zm5.1 2.72a1.42 1.42 0 0 0-1.73 1l-.9 3.45a1.43 1.43 0 1 0 2.76.72l.89-3.45a1.42 1.42 0 0 0-1.02-1.72Z"/>`),
		g.Group(children),
	)
}