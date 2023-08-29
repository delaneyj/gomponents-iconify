package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticSquareBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464Z" opacity=".5"/><path fill-rule="evenodd" d="M16.651 13.86a4.605 4.605 0 1 0-7.715-2.106a.72.72 0 0 1-.172.692L6.289 14.92a.987.987 0 0 0-.283.807l.155 1.392a.658.658 0 0 0 .188.393l.14.139a.657.657 0 0 0 .392.188l1.392.155a.987.987 0 0 0 .807-.283l.296-.297l-1.163-1.15a.75.75 0 0 1 1.055-1.066l1.166 1.153l.003.003l1.118-1.118a.72.72 0 0 1 .69-.172a4.602 4.602 0 0 0 4.406-1.204Zm-4.26-4.136a1.333 1.333 0 1 1 1.885 1.885a1.333 1.333 0 0 1-1.886-1.885Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}