package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kvsitso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 21.468L21.468 5.5m5.161 0L42.5 21.468m-34.274-.097V42.5H39.58V21.371"/><rect width="6.86" height="6.86" x="11.44" y="23.805" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.46"/><rect width="6.86" height="6.86" x="20.495" y="23.805" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.46"/><rect width="6.86" height="6.86" x="29.55" y="23.805" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.46"/><rect width="6.86" height="6.86" x="11.44" y="32.763" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.46"/><rect width="6.86" height="6.86" x="20.495" y="32.763" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.46"/><rect width="6.86" height="6.86" x="29.55" y="32.763" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.46"/>`),
		g.Group(children),
	)
}