package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Momox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.725" cy="28.204" r="8.775" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.725" cy="28.204" r="3.525" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.725" cy="28.204" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 31.189s6.3-.3 11.025 2.85v-20.16C13.8 10.729 7.5 11.029 7.5 11.029Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 14.104h-3v19.2h12.782m9.246-1.966a18.544 18.544 0 0 0-8.003 2.7m0-20.159c4.725-3.15 11.025-2.85 11.025-2.85v10.09m-1.966 12.185h-7.816m9.782-19.2h3v5.598"/>`),
		g.Group(children),
	)
}