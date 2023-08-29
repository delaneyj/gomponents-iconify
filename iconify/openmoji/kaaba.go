package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaaba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#3F3F3F" d="M16.5 16h39v39h-39z"/><path fill="#F1B31C" d="M16.5 23h39v4h-39z"/><path fill="#9B9B9A" d="M22.5 47h5v8h-5z"/><path fill="#F1B31C" d="m36 31l-2 3.5l2 3.5l2-3.5z"/><circle cx="25" cy="34.5" r="2.5" fill="#F1B31C"/><circle cx="46.5" cy="34.5" r="2.5" fill="#F1B31C"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M16.5 16h39v39h-39z"/><path d="M16.5 23h39v4h-39zm6 24h5v8h-5zM36 31l-2 3.5l2 3.5l2-3.5z"/><circle cx="25" cy="34.5" r="2.5"/><circle cx="46.5" cy="34.5" r="2.5"/></g>`),
		g.Group(children),
	)
}