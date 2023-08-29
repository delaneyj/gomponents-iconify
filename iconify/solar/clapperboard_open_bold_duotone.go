package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardOpenBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 11h12c1.886 0 2.828 0 3.414.586C20 12.172 20 13.114 20 15v1c0 2.828 0 4.243-.879 5.121C18.243 22 16.828 22 14 22h-4c-2.828 0-4.243 0-5.121-.879C4 20.243 4 18.828 4 16v-5Z" opacity=".5"/><path d="m13.318 8.2l.6-5.034l-4.71 1.262l-.042.285l-.604 5.064l4.719-1.264l.037-.313Zm-9.835-.685c-.208.776.034 1.68.518 3.485L7 10.196l.039-.314l.6-5.034l-.103.028c-1.805.484-2.708.726-3.276 1.294a3 3 0 0 0-.777 1.345Zm14.765-.333l-3.407.913l.604-5.065l.055-.272a2.62 2.62 0 0 1 .344-.048a3 3 0 0 1 2.887 1.666c.13.265.22.602.401 1.275c.06.225.09.337.097.435a1 1 0 0 1-.556.962c-.088.044-.2.074-.425.134Z"/></g>`),
		g.Group(children),
	)
}