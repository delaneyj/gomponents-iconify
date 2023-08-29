package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="35.905" cy="36.014" r="27.035" fill="#fcea2b"/><circle cx="36.006" cy="36.037" r="21.871" fill="#fff"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M55.11 25.38a21.863 21.863 0 1 1-8.095-8.245"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M62.94 35.997a27.046 27.046 0 1 1-5.266-16.038"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m47.394 21.578l11.038-1.16l-1.16-11.038m-7.297 26.974H35.891V18.52m0 35.391v-3.845M21.143 36.354h-3.057h0"/><circle cx="35.891" cy="36.354" r="3.737"/><circle cx="48.694" cy="47.937" r="1.48"/><circle cx="23.087" cy="24.717" r="1.48"/><circle cx="23.087" cy="47.937" r="1.48"/>`),
		g.Group(children),
	)
}