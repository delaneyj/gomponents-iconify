package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgressFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M559 197H441V-68h118v265zm0 735H441V667h118v265zm176-441V373h265v118H735zM292 307L105 120l83-83l187 187zm520 520L625 640l83-83l187 187zm0-790l83 83l-187 187l-83-83zM188 827l-83-83l187-187l83 83z"/>`),
		g.Group(children),
	)
}