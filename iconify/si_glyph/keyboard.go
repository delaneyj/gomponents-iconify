package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 4.031V11h14.997V4.031H1zM11.969 5h1.047v1.023h-1.047V5zm1.047 1.969v1.047h-1.031V6.969h1.031zm-3.049-1.99h1.05v1.06h-1.05v-1.06zm1.055 1.996v1.062H9.979V6.975h1.043zm-3.04-1.996h1.033v1.036H7.982V4.979zm-.017 1.996h1.078v1.084H7.965V6.975zm-2.993-1.98h1.054v1.028H4.972V4.995zM7 6.988v1.049H5.985V6.988H7zm-1.969-.013v1.062H3.97V6.975h1.061zM2.969 4.984h1.062v1.047H2.969V4.984zm-1 1.995H3v1.037H1.969V6.979zM4 10H2V9h2v1zm8.021.021H4.969V8.968h7.052v1.053zM15 10h-2V9h2v1zm.016-2h-1.031V6.969h1.031V8zm0-1.977h-1.047V5h1.047v1.023z"/>`),
		g.Group(children),
	)
}