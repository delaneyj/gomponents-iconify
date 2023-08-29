package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOfTheHornsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M13 45.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><path fill="#ffe500" d="M30 22.08L31.93 9a2.5 2.5 0 1 1 5 .69C34.67 26.25 34.69 23.16 34.69 32a10 10 0 0 1-10.63 10a10.2 10.2 0 0 1-9.31-10.29v-7.17a11.57 11.57 0 0 0-.58-3.65L12 14.27a2.28 2.28 0 0 1 4.29-1.56l3.78 9.37Z"/><path fill="#fff48c" d="M34 9.57a2.88 2.88 0 0 1 2.64 1.72c.07-.51.14-1 .21-1.6a2.5 2.5 0 1 0-5-.69l-.26 1.81A2.86 2.86 0 0 1 34 9.57ZM12.61 16.1a2.45 2.45 0 0 1 4.51-1.33l-.83-2.06A2.28 2.28 0 0 0 12 14.27l.61 1.86Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30 22.08L31.93 9a2.5 2.5 0 1 1 5 .69C34.67 26.25 34.69 23.16 34.69 32a10 10 0 0 1-10.63 10a10.2 10.2 0 0 1-9.31-10.29v-7.17a11.57 11.57 0 0 0-.58-3.65L12 14.27a2.28 2.28 0 0 1 4.29-1.56l3.78 9.37Z"/><path fill="#ffe500" d="M20.07 19.43h5.32v10.96h-5.32Z"/><path fill="#fff48c" d="M22.73 19.43a2.66 2.66 0 0 0-2.66 2.65v2.4a2.66 2.66 0 0 1 2.66-2.65a2.66 2.66 0 0 1 2.66 2.65v-2.4a2.66 2.66 0 0 0-2.66-2.65Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M20.07 19.43h5.32v10.96h-5.32Z"/><path fill="#ffe500" d="M25.39 19.43h5.32v10.96h-5.32Z"/><path fill="#fff48c" d="M28 19.43a2.65 2.65 0 0 0-2.65 2.65v2.4A2.65 2.65 0 0 1 28 21.83a2.66 2.66 0 0 1 2.66 2.65v-2.4A2.66 2.66 0 0 0 28 19.43Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.39 19.43h5.32v10.96h-5.32Z"/><path fill="#ffe500" d="M34.69 31.75a5.18 5.18 0 0 0-5.19-5.18h-8.07a1.37 1.37 0 0 0-1.38 1.37a3.81 3.81 0 0 0 3.81 3.81Z"/><path fill="#fff48c" d="M34.69 31.75a5.18 5.18 0 0 0-5.19-5.18h-8.07a1.37 1.37 0 0 0-1.38 1.37a3.71 3.71 0 0 0 .24 1.3a1.41 1.41 0 0 1 1.14-.59H30a5.18 5.18 0 0 1 4.69 3.1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.5 37.5v-1.31a4.44 4.44 0 0 1 4.44-4.44h-4.08a3.81 3.81 0 0 1-3.81-3.81a1.37 1.37 0 0 1 1.38-1.37h8.07a5.18 5.18 0 0 1 5.19 5.18"/>`),
		g.Group(children),
	)
}