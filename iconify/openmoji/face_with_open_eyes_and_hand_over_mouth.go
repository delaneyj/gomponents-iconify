package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithOpenEyesAndHandOverMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#fcea2b"><path d="M36 12.242A23.758 23.758 0 1 0 59.758 36A23.785 23.785 0 0 0 36 12.242Z"/><path d="M31 44.6c1.591.854.486 3.626.5 4.7a26.463 26.463 0 0 0 6.9-4c1.595.047 7.958 1.281 5.8 3.4c0 0 1.2.4 0 3c-.7.5-.2 2.4-1 3.4c-4.136 4.135-11.66 3.291-14 7.474l-5.2-5.2c3.69-3.36 2.558-10.133 7-12.774Z"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.062 54.297a22.996 22.996 0 1 1 22.448 3.077"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.07 28.743a7.174 7.174 0 0 1 4.91-1.636a7.078 7.078 0 0 1 4.09 1.636m9.86 0a7.174 7.174 0 0 1 4.908-1.636a7.078 7.078 0 0 1 4.091 1.636"/><path d="M30 35a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3m18 0a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m29.6 61.8l1.8-1.8c2.5-.6 9.9-2.8 11.5-5.5a1.542 1.542 0 0 0-.1-2.1c-.1-.1-.1-.2-.2-.2c.5-.3.8-.6.9-.6a2.108 2.108 0 0 0 .7-2.9a1.205 1.205 0 0 0-.5-.5a1.807 1.807 0 0 0-.1-2.1c-.5-.6-1.6-1.1-3.1 0a.779.779 0 0 0-.3-.7a1.905 1.905 0 0 0-2.7-.4a.098.098 0 0 0-.1.1a31.955 31.955 0 0 1-6.3 4c.503-1.286 1.69-5.814-.6-5.1c-.6.2-1.2 1.7-1.5 2.2c-1.738 2.511-2.258 5.218-2.8 8.1c.031 1.077-1.076 1.676-1.8 2.4"/>`),
		g.Group(children),
	)
}