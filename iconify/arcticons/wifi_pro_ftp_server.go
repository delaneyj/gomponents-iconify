package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiProFtpServer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.68 13.015H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9a2.18 2.18 0 0 0-2.18 2.18v7.34h39v-3.42a1.83 1.83 0 0 0-1.79-1.87Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 18.295h-39v18.72a2.18 2.18 0 0 0 2.16 2.2h34.66a2.18 2.18 0 0 0 2.18-2.18v-.02Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.534 30.484a3.939 3.939 0 0 0-5.032-.014v.014m2.506 1.612a1.58 1.58 0 1 0 1.58 1.58h0a1.58 1.58 0 0 0-1.58-1.58Zm4.928-4.343a7.744 7.744 0 0 0-9.856 0m12.224-2.866a11.45 11.45 0 0 0-14.608 0"/>`),
		g.Group(children),
	)
}