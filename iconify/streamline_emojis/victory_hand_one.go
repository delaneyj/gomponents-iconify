package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VictoryHandOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M13 45.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><path fill="#ffe500" d="M34.13 7a2.5 2.5 0 0 0-2.82 2.13l-2 13.67l-2-13.63a2.5 2.5 0 1 0-4.95.69l1.45 10.34a2.65 2.65 0 0 0-4.39 2v2.4a2.66 2.66 0 0 0-5.31 0v7.21a10.19 10.19 0 0 0 9.3 10.29a10 10 0 0 0 10.66-10.2c0-8.53 0-5.68 2.19-22.06A2.5 2.5 0 0 0 34.13 7Z"/><path fill="#fff48c" d="M22.73 12.38a2.6 2.6 0 0 1 4.94-.85l-.34-2.34a2.5 2.5 0 1 0-4.95.69ZM34.13 7a2.5 2.5 0 0 0-2.82 2.13L31 11.27a2.61 2.61 0 0 1 5 .87c.1-.72.2-1.48.31-2.3A2.5 2.5 0 0 0 34.13 7Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.32 31.9h0a4.44 4.44 0 0 0-4.44 4.44v1.31"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.13 7a2.5 2.5 0 0 0-2.82 2.13l-2 13.67l-2-13.63a2.5 2.5 0 1 0-4.95.69l1.45 10.34a2.65 2.65 0 0 0-4.39 2v2.4a2.66 2.66 0 0 0-5.31 0v7.21a10.19 10.19 0 0 0 9.3 10.29a10 10 0 0 0 10.66-10.2c0-8.53 0-5.68 2.19-22.06A2.5 2.5 0 0 0 34.13 7Z"/><path fill="#ffe500" d="M14.13 21.98h5.32v8.56h-5.32Z"/><path fill="#ffe500" d="M19.44 19.57h5.32v10.96h-5.32Z"/><path fill="#fff48c" d="M16.79 22a2.66 2.66 0 0 0-2.66 2.65V27a2.66 2.66 0 0 1 2.66-2.65A2.65 2.65 0 0 1 19.44 27v-2.4a2.65 2.65 0 0 0-2.65-2.6Zm5.31-2.43a2.67 2.67 0 0 0-2.66 2.66v2.4A2.66 2.66 0 0 1 22.1 22a2.66 2.66 0 0 1 2.66 2.65v-2.4a2.67 2.67 0 0 0-2.66-2.68Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.13 21.98h5.32v8.56h-5.32Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.44 19.57h5.32v10.96h-5.32Z"/><path fill="#ffe500" d="M34.07 31.9a5.2 5.2 0 0 0-5.19-5.19H20.8a1.37 1.37 0 0 0-1.37 1.38a3.81 3.81 0 0 0 3.81 3.81h4.08"/><path fill="#fff48c" d="M34.07 31.9a5.2 5.2 0 0 0-5.19-5.19H20.8a1.37 1.37 0 0 0-1.37 1.38a3.75 3.75 0 0 0 .24 1.3a1.37 1.37 0 0 1 1.13-.59h8.52a5.21 5.21 0 0 1 4.75 3.1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.07 31.9h0a5.2 5.2 0 0 0-5.19-5.19H20.8a1.37 1.37 0 0 0-1.37 1.38h0a3.81 3.81 0 0 0 3.81 3.81h4.08"/>`),
		g.Group(children),
	)
}