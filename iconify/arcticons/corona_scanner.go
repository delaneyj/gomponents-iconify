package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoronaScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 20.473V9.5h-29v10.973a3.527 3.527 0 0 1 0 7.054V38.5h29V27.527a3.527 3.527 0 0 1 0-7.054Z"/><rect width="7.054" height="7.054" x="14.986" y="14.986" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".784"/><rect width="7.054" height="7.054" x="14.986" y="25.959" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".784"/><rect width="7.054" height="7.054" x="25.959" y="14.986" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".784"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.311 28.311h2.351v2.351h-2.351zm2.351-2.352h2.351v2.351h-2.351zm0 4.703h2.351v2.351h-2.351zm-4.703-4.703h2.351v2.351h-2.351zm0 4.703h2.351v2.351h-2.351zM12 5.5H5.5V12m37 0V5.5H36m0 37h6.5V36m-37 0v6.5H12"/>`),
		g.Group(children),
	)
}