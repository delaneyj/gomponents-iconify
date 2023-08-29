package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Privacyshade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.427 21.096a9.548 9.548 0 1 1-19.095 0a9.548 9.548 0 1 1 19.095 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 21.329c.807-6.712-5.48-12.24-12.12-12.24s-12.928 5.528-12.121 12.24c.82 6.826 7.495 16.6 12.12 16.442c4.626.158 11.3-9.616 12.121-16.442Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.053 31.161h6.766v7.578H9.18v-7.578h6.524m20.361-10.9h2.752v7.578H33.93m-20.102 0H9.181v-7.578h2.512m14.215-11h12.91v7.578h-3.345m-23.186 0H9.181V9.261h12.67"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}