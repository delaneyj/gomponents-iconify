package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomansSandal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M3.08 35.53a21.71 1.97 0 1 0 43.42 0a21.71 1.97 0 1 0-43.42 0Z" opacity=".15"/><path fill="#f7e5c6" d="M2.75 21.3L5 34.17A1 1 0 0 0 6 35h6.61a1 1 0 0 0 1-1V21.3Z"/><path fill="#f0d5a8" d="M3.24 24H5.9a15.38 15.38 0 0 1 7.71 2.54V21.3H2.75Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2.75 21.3L5 34.17A1 1 0 0 0 6 35h6.61a1 1 0 0 0 1-1V21.3Z"/><path fill="#debb7e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24.08 31a20.45 20.45 0 0 0 12.18 4h7.62a.49.49 0 0 0 .45-.29l1.3-2.84a.5.5 0 0 0-.46-.71H38a19.11 19.11 0 0 1-11-3.52l-1.9-1.33c-4.81-3.41-8.48-8.69-16.44-9.18H3.25a.5.5 0 0 0-.5.5v3.67h4c6.25 0 12.31 5.95 17.33 9.7Z"/><path fill="#f7e5c6" d="m22.75 30l1.33 1a20.45 20.45 0 0 0 12.18 4h1.05c1.46-3.64 2.48-10.19 2.85-12.81a1.2 1.2 0 0 0-.9-1.32A31.64 31.64 0 0 1 28 16.09a.7.7 0 0 0-1.08.6a20 20 0 0 1-1.44 8.11a29.09 29.09 0 0 1-2.73 5.2Z"/><path fill="#fff5e3" d="M26.47 21.43A49.46 49.46 0 0 0 39 25a1.69 1.69 0 0 1 .71.29c.22-1.27.38-2.37.49-3.14a1.2 1.2 0 0 0-.9-1.32A31.64 31.64 0 0 1 28 16.09a.7.7 0 0 0-1.08.6a19.9 19.9 0 0 1-.45 4.74Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m22.75 30l1.33 1a20.45 20.45 0 0 0 12.18 4h1.05c1.46-3.64 2.48-10.19 2.85-12.81a1.2 1.2 0 0 0-.9-1.32A31.64 31.64 0 0 1 28 16.09a.7.7 0 0 0-1.08.6a20 20 0 0 1-1.44 8.11a29.09 29.09 0 0 1-2.73 5.2Z"/>`),
		g.Group(children),
	)
}