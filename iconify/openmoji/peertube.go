package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peertube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="m16.757 11.706l19.11 11.793l-18.594 11.449l-.516-23.242z"/><path fill="#e27022" d="M36.556 47.688L55.838 35.55L36.642 24.015l-.086 23.673z"/><path fill="#9b9b9a" d="m17.101 35.895l19.196 12.309l-19.282 12.052l.086-24.361z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36.263 23.347l-19.48-12.389v24.778l19.48-12.389zm19.48 12.389l-19.48-12.389v24.778l19.48-12.389zm-19.48 12.389l-19.48-12.389v24.778l19.48-12.389z"/>`),
		g.Group(children),
	)
}