package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easypaisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M24.602 4.5c8.516 0 15.42 5.717 15.42 12.77s-6.902 11.344-15.42 11.344q-12.133-.245-17.19-7.642a1.647 1.647 0 0 1-.246-1.242Q10.057 5.298 24.602 4.5Zm-.575 7.944q-7.345.744-8.989 6.993q2.24 1.975 8.99 2.192c4.52-.07 7.259-1.389 7.312-4.452c-.281-3.384-3.571-4.861-7.313-4.733Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M6.396 24.71c2.526 7.247 9.492 11.28 16.654 11.28c5.94 0 9.876-2.57 11.94-6.45l6.656 4.17c-2.32 5.675-9.198 9.79-17.323 9.79c-10.029 0-18.617-6.76-17.93-18.511c0-.094.001-.187.003-.279Z"/>`),
		g.Group(children),
	)
}