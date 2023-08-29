package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.586 12.336C2 12.922 2 13.864 2 15.75c0 1.886 0 2.828.586 3.414c.586.586 1.528.586 3.414.586h12c1.886 0 2.828 0 3.414-.586c.586-.586.586-1.528.586-3.414c0-1.886 0-2.828-.586-3.414c-.586-.586-1.528-.586-3.414-.586H6c-1.886 0-2.828 0-3.414.586ZM6 16.75a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm4-.75a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Z" clip-rule="evenodd"/><path d="M3.651 4.378a.75.75 0 0 0-1.302.744l3.787 6.628H7.86l-.209-.372l-4-7Zm16.699 0a.75.75 0 1 1 1.301.744l-3.787 6.628h-1.723l.208-.372l4-7Z" opacity=".5"/><path fill-rule="evenodd" d="M12.084 3.5a4.752 4.752 0 0 0-4.39 2.934a.75.75 0 1 1-1.387-.574a6.252 6.252 0 0 1 11.553 0a.75.75 0 0 1-1.386.574a4.752 4.752 0 0 0-4.39-2.934Z" clip-rule="evenodd" opacity=".4"/><path fill-rule="evenodd" d="M12.085 6a2.251 2.251 0 0 0-2.16 1.618a.75.75 0 0 1-1.44-.42a3.751 3.751 0 0 1 7.2 0a.75.75 0 0 1-1.44.42A2.251 2.251 0 0 0 12.085 6Z" clip-rule="evenodd" opacity=".7"/><path d="M13.084 7.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}