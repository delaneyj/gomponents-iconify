package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardiogramEOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M12 29a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H12Zm-1 6a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Z"/><path fill-rule="evenodd" d="M12 9a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3h3a3 3 0 0 1 3 3v27a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V12a3 3 0 0 1 3-3h3ZM8 23h6.171l1.47-2.466l1.815 6.661L21.49 22h4.278a1 1 0 0 0 0-2H20.51l-2.177 2.805l-2-7.339L13.035 21H8v-9a1 1 0 0 1 1-1h3a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3h3a1 1 0 0 1 1 1v27a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V23ZM24 9v2a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Zm15 5a3 3 0 0 0-3 3v20.303l3 4.5l3-4.5V17a3 3 0 0 0-3-3Zm-1 3a1 1 0 1 1 2 0v2h-2v-2Zm0 19.697V21h2v15.697l-1 1.5l-1-1.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}