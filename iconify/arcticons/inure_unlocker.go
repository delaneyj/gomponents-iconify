package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InureUnlocker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.5 32.126l8.277-14.238m-7.933 25.206l10.978-18.885M24 43.5c10.819 0 14.91-7.14 14.91-14.902v-3.737c0-5.606-2.426-6.973-8.823-6.973H17.913c-6.397 0-8.822 1.367-8.822 6.973v3.737C9.09 36.361 13.18 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.95 17.942V12.55a8.05 8.05 0 0 1 15.585-2.84m.515 8.232v-3.123"/>`),
		g.Group(children),
	)
}