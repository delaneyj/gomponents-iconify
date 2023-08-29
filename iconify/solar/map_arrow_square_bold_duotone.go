package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapArrowSquareBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" clip-rule="evenodd" opacity=".5"/><path d="m13.423 17.362l3.512-9.166a.863.863 0 0 0-1.131-1.13l-9.166 3.511c-.83.319-.857 1.483-.04 1.731l3.477 1.057c.27.082.478.29.56.56l1.057 3.477c.248.817 1.412.79 1.73-.04Z"/></g>`),
		g.Group(children),
	)
}